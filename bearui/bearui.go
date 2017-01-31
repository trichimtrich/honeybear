package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"

	"strings"

	elastic "gopkg.in/olivere/elastic.v5"
)

var hostELK = "http://192.168.200.1:9200"
var hostDocker = "tcp://192.168.200.106:4243"
var hostGlanceFarm = "http://192.168.200.106:61208/api/2"
var hostGlanceProxy = "http://192.168.200.105:61208/api/2"
var hostEngine = "http://192.168.200.1:6969"
var listenHost = "0.0.0.0"
var listenPort = "9090"
var netClient = &http.Client{
	Timeout: time.Second * 10,
}
var Docker *client.Client
var elkClient *elastic.Client

type UserID struct {
	IP          string `json:"ip"`
	UserAgent   string `json:"agent"`
	ProductSess string `json:"product"`
	FarmSess    string `json:"farm"`
	Time        string `json:"time"`
	Status      string `json:"status"`
}

type UserHistory struct {
	UID         string `json:"uid"`
	IP          string `json:"ip"`
	UserAgent   string `json:"agent"`
	ProductSess string `json:"product"`
	FarmSess    string `json:"farm"`
	Time        string `json:"time"`
	Status      string `json:"status"`
}

type Request struct {
	UID       string                   `json:"uid"`
	URL       string                   `json:"url"`
	UserAgent string                   `json:"agent"`
	Cookie    []map[string]interface{} `json:"cookie"`
	GetForm   map[string][]string      `json:"get"`
	PostForm  map[string][]string      `json:"post"`
	Time      string                   `json:"time"`
}

type Query struct {
	RID   string `json:"rid"`
	Query string `json:"query"`
	Time  string `json:"time"`
}

type Command struct {
	RID  string `json:"rid"`
	Cmd  string `json:"cmd"`
	Time string `json:"time"`
}

type Snapshot struct {
	Name        string `json:"name"`
	Description string `json:"desc"`
	LearnStart  string `json:"learnStart"`
	LearnEnd    string `json:"learnEnd"`
	VerifyStart string `json:"verifyStart"`
	VerifyEnd   string `json:"verifyEnd"`
	Time        string `json:"time"`
}

type RequestCheck struct {
	Time        string            `json:"time"`
	UID         string            `json:"uid"`
	IP          string            `json:"ip"`
	Status      string            `json:"status"`
	RID         string            `json:"rid"`
	URL         string            `json:"url"`
	Param       map[string]string `json:"param"`
	Violation   string            `json:"violation"`
	Query       []string          `json:"query"`
	Cmd         []string          `json:"cmd"`
	WebObject   map[string]string `json:"webobj"`
	QueryObject [][]string        `json:"queryobj"`
	CmdObject   [][]string        `json:"cmdobj"`
}

func func404(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/404.tpl")
	t.Execute(w, nil)
}

func funcIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		func404(w, r)
	} else {
		_, err := r.Cookie("bear")
		if err != nil {
			t, _ := template.ParseFiles("template/login.tpl")
			t.Execute(w, nil)
		} else {
			//bearboard
			queryCount, err := elkClient.Count().Index("request").Type("check").Query(elastic.NewMatchQuery("violation", "query")).Do(context.Background())
			if err != nil {
				fmt.Println("[Index] Cannot get sqli count:", err)
			}

			cmdCount, err := elkClient.Count().Index("request").Type("check").Query(elastic.NewMatchQuery("violation", "cmd")).Do(context.Background())
			if err != nil {
				fmt.Println("[Index] Cannot get cmdi count:", err)
			}

			webCount, err := elkClient.Count().Index("request").Type("check").Query(elastic.NewMatchQuery("violation", "web")).Do(context.Background())
			if err != nil {
				fmt.Println("[Index] Cannot get webobject count:", err)
			}

			//get userlist
			hackerCount := 0
			suspectCount := 0
			listUID := make([]map[string]string, 0)
			uidELK, err := elkClient.Search().Index("uid").Type("bear").Size(10000).Sort("time", false).Do(context.Background())
			if err != nil {
				fmt.Println("[Index] Cannot open UID ELK:", err)
			} else {
				fmt.Println("[UID] Total hit:", uidELK.TotalHits(), "- Total time ms:", uidELK.TookInMillis)
				for i, item := range uidELK.Hits.Hits {
					var curUID UserID
					err = json.Unmarshal(*item.Source, &curUID)
					if err != nil {
						fmt.Println("[Index] UID Fail:", i, item.Id)
					} else {
						if curUID.Status == "1" {
							suspectCount++
						}
						if curUID.Status == "2" {
							hackerCount++
						}
						listUID = append(listUID, map[string]string{
							"uid":    item.Id,
							"ip":     curUID.IP,
							"agent":  curUID.UserAgent,
							"status": curUID.Status,
						})
					}
				}
			}
			t, _ := template.ParseFiles("template/index.tpl")
			t.Execute(w, map[string]interface{}{
				"stat": []int{hackerCount, suspectCount, int(queryCount), int(cmdCount), int(webCount)},
				"user": listUID,
			})
		}
	}
}

//--------------------------------bearfarm
func funcHost(w http.ResponseWriter, r *http.Request) {
	hostStat := make(map[string]interface{})
	var jsonObj interface{}

	resp, err := netClient.Get(hostGlanceFarm + "/all")
	if err != nil {
		log.Println("Cannot handle api /all: ", err)
	} else {
		buf, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(buf, &jsonObj)
		if err != nil {
			log.Println("Cannot decode api /all: ", err)
		} else {
			curObj := jsonObj.(map[string]interface{})

			//uptime, time
			hostStat["uptime"] = curObj["uptime"].(string)
			hostStat["now"] = curObj["now"].(string)

			//hostname, kernel, os
			pluginSystem := curObj["system"].(map[string]interface{})
			hostStat["hostname"] = pluginSystem["hostname"].(string)
			hostStat["kernel"] = pluginSystem["os_version"].(string)
			hostStat["os"] = fmt.Sprintf("%s - %s %s", pluginSystem["os_name"].(string), pluginSystem["linux_distro"].(string), pluginSystem["platform"].(string))

			//ip
			hostStat["ip"] = curObj["ip"].(map[string]interface{})["address"].(string)

			//cpuinfo - quicklook
			pluginQuickLook := curObj["quicklook"].(map[string]interface{})
			hostStat["cpu1"] = pluginQuickLook["cpu_name"].(string)
			hostStat["cpu2"] = fmt.Sprintf("%dx%.2f GHz", len(pluginQuickLook["percpu"].([]interface{})), pluginQuickLook["cpu_hz_current"].(float64)/1000000000)

			//logical memory
			hostStat["memory"] = fmt.Sprintf("%.f MB", curObj["mem"].(map[string]interface{})["total"].(float64)/float64(1024*1024))

			//swap memory
			hostStat["swap"] = fmt.Sprintf("%.f MB", curObj["memswap"].(map[string]interface{})["total"].(float64)/float64(1024*1024))

			//file system
			pluginFS := curObj["fs"].([]interface{})[0].(map[string]interface{})
			hostStat["storage"] = fmt.Sprintf("Used %.2f GiB / Total %.2f GiB", pluginFS["used"].(float64)/float64(1024*1024*1024), pluginFS["size"].(float64)/float64(1024*1024*1024))

			//Processlist
			hostStat["process"] = curObj["processlist"]

			//Network
			hostStat["network"] = curObj["network"]

			//Container
			hostStat["container"] = curObj["docker"]
		}
		defer resp.Body.Close()
	}

	hostStat["glance"] = hostGlanceFarm
	hostStat["hihi"] = map[string]string{"zz": "yo", "yy": "lo"}

	t, _ := template.ParseFiles("template/farm/host.tpl")
	t.Execute(w, hostStat)
}

func funcHostStat(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	p := r.FormValue("p")
	resp, _ := netClient.Get(hostGlanceFarm + "/" + p)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	w.Header()["Content-type"] = []string{"application/json"}
	w.Write(body)
}

func funcContainer(w http.ResponseWriter, r *http.Request) {
	listContainer := make(map[string]interface{})
	var onlyRunning = false
	r.ParseForm()
	if r.Method == "GET" && r.FormValue("full") == "false" {
		onlyRunning = true
		listContainer["full1"] = "onclick=\"document.location='/farm/container'\""
		listContainer["full2"] = ""
	} else {
		listContainer["full1"] = "onclick=\"document.location='/farm/container?full=false'\""
		listContainer["full2"] = "checked"
	}

	containers, _ := Docker.ContainerList(context.Background(), types.ContainerListOptions{
		All: true,
	})

	a, _ := Docker.Info(context.Background())

	arr := make([]map[string]string, 0)
	c := 0
	for _, container := range containers {
		if onlyRunning && container.State != "running" {
			continue
		}
		d := make(map[string]string)
		d["id"] = container.ID[:12]
		d["name"] = container.Names[0]
		if val, ok := container.NetworkSettings.Networks["bridge"]; ok {
			d["ip"] = val.IPAddress
		} else {
			d["ip"] = "N/A"
		}
		d["image"] = container.Image
		d["command"] = container.Command
		d["state"] = container.State
		d["status"] = container.Status
		d["hostname"] = a.Name
		if container.State == "running" {
			c++
		}
		arr = append(arr, d)
	}

	listContainer["container"] = arr
	listContainer["running"] = c
	listContainer["total"] = len(containers)

	t, _ := template.ParseFiles("template/farm/container.tpl")
	t.Execute(w, listContainer)
}

func funcContainerStat(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")
	listStat := make(map[string]interface{})

	info, err := Docker.ContainerInspect(context.Background(), id)
	if err != nil {
		func404(w, r)
		return
	} else {
		listStat["name"] = info.Name
		a, _ := Docker.Info(context.Background())
		listStat["host"] = a.Name
		listStat["hostname"] = info.Config.Hostname
		listStat["pid"] = info.State.Pid
		listStat["id"] = info.ID
		listStat["image"] = info.Config.Image
		listStat["state"] = info.State.Status
		if val, ok := info.NetworkSettings.Networks["bridge"]; ok {
			listStat["ip"] = val.IPAddress
		} else {
			listStat["ip"] = "N/A"
		}
	}
	t, _ := template.ParseFiles("template/farm/containerstat.tpl")
	t.Execute(w, listStat)
}

func funcContainerLiveStat(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")
	listStat := make(map[string]interface{})

	stat, err := Docker.ContainerStats(context.Background(), id, false)
	defer stat.Body.Close()
	if err != nil {
		w.Write([]byte("fail"))
	} else {
		buf, _ := ioutil.ReadAll(stat.Body)
		var jsonObj interface{}
		err = json.Unmarshal(buf, &jsonObj)
		if err != nil {
			w.Write([]byte("fail"))
		} else {
			curObj := jsonObj.(map[string]interface{})
			listStat["cpu"] = curObj["cpu_stats"].(map[string]interface{})["cpu_usage"].(map[string]interface{})["total_usage"].(float64) / float64(1024*1024)
			listStat["mem"] = curObj["memory_stats"].(map[string]interface{})["usage"].(float64) / float64(1024*1024)
			listStat["net"] = curObj["networks"]
			listStat["disk"] = curObj["storage_stats"]
			buf2, _ := json.Marshal(listStat)
			w.Header()["Content-type"] = []string{"application/json"}
			w.Write(buf2)
		}
	}
}

func funcContainerControl(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	rid := r.FormValue("id")
	con := r.FormValue("do")
	if rid != "" {
		if con == "start" {
			err := Docker.ContainerStart(context.Background(), rid, types.ContainerStartOptions{})
			w.Write([]byte(fmt.Sprint(err)))
		} else if con == "restart" {
			timeout := 10 * time.Second
			err := Docker.ContainerRestart(context.Background(), rid, &timeout)
			w.Write([]byte(fmt.Sprint(err)))
		} else if con == "stop" {
			timeout := 10 * time.Second
			err := Docker.ContainerStop(context.Background(), rid, &timeout)
			w.Write([]byte(fmt.Sprint(err)))
		} else {
			w.Write([]byte("No Valid Control"))
		}
	} else {
		w.Write([]byte("No RID"))
	}
}

func funcTopology(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/farm/topology.tpl")
	t.Execute(w, nil)
}

//BearProxy
func funcProxy(w http.ResponseWriter, r *http.Request) {
	hostStat := make(map[string]interface{})
	var jsonObj interface{}

	resp, err := netClient.Get(hostGlanceProxy + "/all")
	if err != nil {
		log.Println("Cannot handle api /all: ", err)
	} else {
		buf, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(buf, &jsonObj)
		if err != nil {
			log.Println("Cannot decode api /all: ", err)
		} else {
			curObj := jsonObj.(map[string]interface{})

			//uptime, time
			hostStat["uptime"] = curObj["uptime"].(string)
			hostStat["now"] = curObj["now"].(string)

			//hostname, kernel, os
			pluginSystem := curObj["system"].(map[string]interface{})
			hostStat["hostname"] = pluginSystem["hostname"].(string)
			hostStat["kernel"] = pluginSystem["os_version"].(string)
			hostStat["os"] = fmt.Sprintf("%s - %s %s", pluginSystem["os_name"].(string), pluginSystem["linux_distro"].(string), pluginSystem["platform"].(string))

			//ip
			hostStat["ip"] = curObj["ip"].(map[string]interface{})["address"].(string)

			//cpuinfo - quicklook
			pluginQuickLook := curObj["quicklook"].(map[string]interface{})
			hostStat["cpu1"] = pluginQuickLook["cpu_name"].(string)
			hostStat["cpu2"] = fmt.Sprintf("%dx%.2f GHz", len(pluginQuickLook["percpu"].([]interface{})), pluginQuickLook["cpu_hz_current"].(float64)/1000000000)

			//logical memory
			hostStat["memory"] = fmt.Sprintf("%.f MB", curObj["mem"].(map[string]interface{})["total"].(float64)/float64(1024*1024))

			//swap memory
			hostStat["swap"] = fmt.Sprintf("%.f MB", curObj["memswap"].(map[string]interface{})["total"].(float64)/float64(1024*1024))

			//file system
			pluginFS := curObj["fs"].([]interface{})[0].(map[string]interface{})
			hostStat["storage"] = fmt.Sprintf("Used %.2f GiB / Total %.2f GiB", pluginFS["used"].(float64)/float64(1024*1024*1024), pluginFS["size"].(float64)/float64(1024*1024*1024))

			//Processlist
			hostStat["process"] = curObj["processlist"]

			//Network
			hostStat["network"] = curObj["network"]

		}
		defer resp.Body.Close()
	}

	hostStat["glance"] = hostGlanceProxy
	hostStat["hihi"] = map[string]string{"zz": "yo", "yy": "lo"}

	t, _ := template.ParseFiles("template/proxy/stat.tpl")
	t.Execute(w, hostStat)
}

func funcProxyStat(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	p := r.FormValue("p")
	resp, _ := netClient.Get(hostGlanceProxy + "/" + p)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	w.Header()["Content-type"] = []string{"application/json"}
	w.Write(body)
}

func funcUID(w http.ResponseWriter, r *http.Request) {
	uidELK, err := elkClient.Search().Index("uid").Type("bear").Sort("time", false).Size(10000).Do(context.Background())
	outUID := make(map[string]interface{})
	outUID["status"] = []int{0, 0, 0, 0}
	if err != nil {
		log.Println("Cannot open UID ELK:", err)
	} else {
		listUID := make([]map[string]string, 0)
		fmt.Println("[UID] Total hit:", uidELK.TotalHits(), "- Total time ms:", uidELK.TookInMillis)
		for i, item := range uidELK.Hits.Hits {
			var curUID UserID
			err = json.Unmarshal(*item.Source, &curUID)
			if err != nil {
				fmt.Println("[UID] Fail:", i, item.Id)
			} else {
				idStatus, _ := strconv.Atoi(curUID.Status)
				outUID["status"].([]int)[idStatus]++
				listUID = append(listUID, map[string]string{
					"uid":     item.Id,
					"ip":      curUID.IP,
					"agent":   curUID.UserAgent,
					"product": curUID.ProductSess,
					"farm":    curUID.FarmSess,
					"time":    curUID.Time,
					"status":  curUID.Status,
				})
			}
		}
		outUID["list"] = listUID
	}
	t, _ := template.ParseFiles("template/proxy/uid.tpl")
	t.Execute(w, outUID)
}

func funcUIDHistory(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	uid := r.FormValue("uid")

	uidELK, err := elkClient.Search().Index("uid").Type("history").Query(elastic.NewMatchQuery("uid", uid)).Sort("time", false).Size(10000).Do(context.Background())
	historyUID := make([]map[string]string, 0)
	if err != nil {
		log.Println("Cannot open HistoryUID ELK:", err)
	} else {
		fmt.Println("[HistoryUID] Total hit:", uidELK.TotalHits(), "- Total time ms:", uidELK.TookInMillis)
		for i, item := range uidELK.Hits.Hits {
			var curUID UserHistory
			err = json.Unmarshal(*item.Source, &curUID)
			if err != nil {
				fmt.Println("[HistoryUID] Fail:", i, item.Id)
			} else {
				historyUID = append(historyUID, map[string]string{
					"uid":     curUID.UID,
					"ip":      curUID.IP,
					"agent":   curUID.UserAgent,
					"product": curUID.ProductSess,
					"farm":    curUID.FarmSess,
					"time":    curUID.Time,
					"status":  curUID.Status,
				})
			}
		}
	}
	t, _ := template.ParseFiles("template/proxy/uidhistory.tpl")
	t.Execute(w, historyUID)
}

func funcTracking(w http.ResponseWriter, r *http.Request) {
	uidELK, err := elkClient.Search().Index("uid").Type("bear").Size(10000).Sort("time", false).Do(context.Background())
	listUID := make([]map[string]string, 0)
	if err != nil {
		log.Println("Cannot open UID ELK for Tracking:", err)
	} else {
		fmt.Println("[UID-Tracking] Total hit:", uidELK.TotalHits(), "- Total time ms:", uidELK.TookInMillis)
		for i, item := range uidELK.Hits.Hits {
			var curUID UserID
			err = json.Unmarshal(*item.Source, &curUID)
			if err != nil {
				fmt.Println("[UID-Tracking] Fail:", i, item.Id)
			} else {
				listUID = append(listUID, map[string]string{
					"uid":     item.Id,
					"ip":      curUID.IP,
					"agent":   curUID.UserAgent,
					"product": curUID.ProductSess,
					"farm":    curUID.FarmSess,
					"time":    curUID.Time,
					"status":  curUID.Status,
				})
			}
		}
	}

	t, _ := template.ParseFiles("template/filter/tracking.tpl")
	t.Execute(w, listUID)
}

func funcTrackingID(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	uid := r.FormValue("uid")

	sdate := r.FormValue("date")
	date, err := time.Parse(time.RFC3339, sdate)
	if err != nil {
		date = time.Now()
	}

	//get nearest day
	reqELK, err := elkClient.Search().Index("request").Type("bear").Query(elastic.NewBoolQuery().Must(elastic.NewRangeQuery("time").Lt(date), elastic.NewMatchQuery("uid", uid))).Sort("time", false).Size(1).Do(context.Background())
	if err != nil {
		log.Println("Cannot open Tracking ELK:", err)
	} else {
		fmt.Println("[UID-Trakcing-Specific] Total hit:", reqELK.TotalHits(), "- Total time ms:", reqELK.TookInMillis)
		if reqELK.TotalHits() > 0 {
			var t Request
			err = json.Unmarshal(*reqELK.Hits.Hits[0].Source, &t)
			if err != nil {
				log.Println("[UID-Trakcing-Specific] Fail: ", err)
			} else {
				tt, _ := time.Parse(time.RFC3339, t.Time)
				fmt.Println(tt)
				y, m, d := tt.Date()
				day1 := time.Date(y, m, d, 0, 0, 0, 0, time.UTC)
				day2 := day1.AddDate(0, 0, 1)
				fmt.Println(day1)
				fmt.Println(day2)

				reqELK, err = elkClient.Search().Index("request").Type("bear").Query(elastic.NewBoolQuery().Must(elastic.NewRangeQuery("time").Lt(day2).Gte(day1), elastic.NewMatchQuery("uid", uid))).Sort("time", false).Size(10000).Do(context.Background())
				fmt.Println("[UID-Trakcing-Specific] Total hit on same day:", reqELK.TotalHits(), "- Total time ms:", reqELK.TookInMillis)
				outTrack := make([]map[string]interface{}, 0)
				outTrack = append(outTrack, map[string]interface{}{
					"next":    day1.Format(time.RFC3339),
					"current": day1.Format("02 Jan 2006"),
				})
				for _, item := range reqELK.Hits.Hits {
					var t2 Request
					err = json.Unmarshal(*item.Source, &t2)
					if err != nil {
						log.Println("[UID-Trakcing-Specific] Fail on id: ", item.Id, err)
					} else {
						tt, _ = time.Parse(time.RFC3339, t2.Time)
						outTrack = append(outTrack, map[string]interface{}{
							"rid":    item.Id,
							"url":    t2.URL,
							"uid":    t2.UID,
							"agent":  t2.UserAgent,
							"cookie": t2.Cookie,
							"get":    t2.GetForm,
							"post":   t2.PostForm,
							"time":   fmt.Sprintf("%.2d:%.2d:%.2d", tt.Hour(), tt.Minute(), tt.Second()),
						})
					}
				}
				outJSON, _ := json.Marshal(outTrack)
				w.Header()["Content-type"] = []string{"application/json"}
				w.Write(outJSON)
				return
			}
		} else {
			fmt.Println("[UID-Trakcing-Specific]", uid, " out of requests")
		}
	}
	w.Write([]byte("fail"))
}

func funcTrackingIDQuery(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	rid := r.FormValue("rid")

	qr, err := elkClient.Search().Index("query").Type("bear").Query(elastic.NewMatchQuery("rid", rid)).Size(10000).Do(context.Background())
	outQuery := make([]map[string]string, 0)
	if err != nil {
		log.Println("Cannot open Query ELK:", err)
	} else {
		for _, item := range qr.Hits.Hits {
			var t Query
			err = json.Unmarshal(*item.Source, &t)
			if err != nil {
				log.Println("[UID-Trakcing-Specific-Query] Fail: ", item.Id, err)
			} else {
				tt, _ := time.Parse(time.RFC3339, t.Time)
				outQuery = append(outQuery, map[string]string{
					"query": t.Query,
					"time":  fmt.Sprintf("%.2d:%.2d:%.2d", tt.Hour(), tt.Minute(), tt.Second()),
				})
			}
		}
		outJSON, _ := json.Marshal(outQuery)
		w.Header()["Content-type"] = []string{"application/json"}
		w.Write(outJSON)
		return
	}
	w.Write([]byte("fail"))
}

func funcTrackingIDCmd(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	rid := r.FormValue("rid")

	qr, err := elkClient.Search().Index("cmd").Type("bear").Query(elastic.NewMatchQuery("rid", rid)).Size(10000).Do(context.Background())
	outQuery := make([]map[string]string, 0)
	if err != nil {
		log.Println("Cannot open CMD ELK:", err)
	} else {
		for _, item := range qr.Hits.Hits {
			var t Command
			err = json.Unmarshal(*item.Source, &t)
			if err != nil {
				log.Println("[UID-Trakcing-Specific-CMD] Fail: ", item.Id, err)
			} else {
				tt, _ := time.Parse(time.RFC3339, t.Time)
				outQuery = append(outQuery, map[string]string{
					"cmd":  t.Cmd,
					"time": fmt.Sprintf("%.2d:%.2d:%.2d", tt.Hour(), tt.Minute(), tt.Second()),
				})
			}
		}
		outJSON, _ := json.Marshal(outQuery)
		w.Header()["Content-type"] = []string{"application/json"}
		w.Write(outJSON)
		return
	}
	w.Write([]byte("fail"))
}

func funcLog(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/log.tpl")
	t.Execute(w, nil)
}

func funcSnapshot(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "POST" && r.FormValue("name") != "" && r.FormValue("desc") != "" && r.FormValue("learn") != "" && r.FormValue("verify") != "" {
		name := r.FormValue("name")
		desc := r.FormValue("desc")
		fmt.Println(name)
		fmt.Println(desc)
		learn := strings.Split(r.FormValue("learn"), " - ")
		verify := strings.Split(r.FormValue("verify"), " - ")
		if len(learn) == 2 && len(verify) == 2 {
			learnStart, err := time.Parse(time.RFC3339, learn[0])
			if err != nil {
				fmt.Println("[Snapshot] Learn[0] parsing failed:", err)
			} else {
				learnEnd, err := time.Parse(time.RFC3339, learn[1])
				if err != nil {
					fmt.Println("[Snapshot] Learn[1] parsing failed:", err)
				} else {
					verifyStart, err := time.Parse(time.RFC3339, verify[0])
					if err != nil {
						fmt.Println("[Snapshot] Verify[0] parsing failed:", err)
					} else {
						verifyEnd, err := time.Parse(time.RFC3339, verify[1])
						if err != nil {
							fmt.Println("[Snapshot] Verify[1] parsing failed:", err)
						} else {
							fmt.Println(learnStart)
							fmt.Println(learnEnd)
							fmt.Println(verifyStart)
							fmt.Println(verifyEnd)
							_, err = elkClient.Index().Index("snapshot").Type("bear").BodyJson(map[string]string{
								"name":        name,
								"desc":        desc,
								"learnStart":  learn[0],
								"learnEnd":    learn[1],
								"verifyStart": verify[0],
								"verifyEnd":   verify[1],
								"time":        time.Now().UTC().Format(time.RFC3339),
							}).Refresh("true").Do(context.Background())
							if err != nil {
								fmt.Println("[Snapshot] Failed to add ELK")
							} else {
								fmt.Println("[Snapshot] Successed!")
							}
						}
					}
				}
			}
		} else {
			fmt.Println("[Snapshot] Learn & verify time wrong format")
		}
	}

	//Current ID
	curResp, _ := netClient.Get(hostEngine + "/current")
	defer curResp.Body.Close()
	curIDs, _ := ioutil.ReadAll(curResp.Body)
	curID := string(curIDs)
	//For snapshot
	sr, err := elkClient.Search().Index("snapshot").Type("bear").Sort("time", false).Size(10000).Do(context.Background())
	outSnap := make([]map[string]string, 0)
	if err != nil {
		fmt.Println("[Snapshot] Fail to get ELK:", err)
	} else {
		for i, item := range sr.Hits.Hits {
			var xx Snapshot
			err = json.Unmarshal(*item.Source, &xx)
			if err != nil {
				fmt.Println("[Snapshot] Cannot parse", i, item.Id, ":", err)
			} else {
				l0, _ := time.Parse(time.RFC3339, xx.LearnStart)
				l1, _ := time.Parse(time.RFC3339, xx.LearnEnd)
				v0, _ := time.Parse(time.RFC3339, xx.VerifyStart)
				v1, _ := time.Parse(time.RFC3339, xx.VerifyEnd)
				t, _ := time.Parse(time.RFC3339, xx.Time)
				outSnap = append(outSnap, map[string]string{
					"id":      item.Id,
					"curID":   curID,
					"name":    xx.Name,
					"desc":    xx.Description,
					"learn0":  l0.Format(time.RFC822),
					"learn1":  l1.Format(time.RFC822),
					"verify0": v0.Format(time.RFC822),
					"verify1": v1.Format(time.RFC822),
					"time":    t.Format(time.RFC822),
				})
			}
		}
	}
	t, _ := template.ParseFiles("template/engine/snapshot.tpl")
	t.Execute(w, outSnap)
}

func funcObject(w http.ResponseWriter, r *http.Request) {
	resp, _ := netClient.Get(hostEngine + "/get")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	var tt interface{}
	_ = json.Unmarshal(body, &tt)

	t, _ := template.ParseFiles("template/engine/object.tpl")
	t.Execute(w, tt)
}

func funcSnapshotSet(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	resp, _ := netClient.Get(hostEngine + "/set?id=" + r.FormValue("id"))
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	w.Write(body)
}

func funcWarning(w http.ResponseWriter, r *http.Request) {
	elkWar, err := elkClient.Search().Index("request").Type("check").Sort("time", false).Size(10000).Do(context.Background())
	countWar := []int{0, 0, 0}
	listWar := make([]RequestCheck, 0)
	if err != nil {
		fmt.Println("[Warning] Cannot get ELK:", err)
	} else {
		for _, item := range elkWar.Hits.Hits {
			var war RequestCheck
			err = json.Unmarshal(*item.Source, &war)
			if err != nil {
				fmt.Println("[Warning] Cannot decode json id", item.Id)
			} else {
				listWar = append(listWar, war)
				if war.Violation == "web" {
					countWar[0]++
				}
				if war.Violation == "query" {
					countWar[1]++
				}
				if war.Violation == "cmd" {
					countWar[2]++
				}
			}
		}
	}
	listOut := make(map[string]interface{})
	listOut["object"] = countWar
	listOut["warning"] = listWar
	t, _ := template.ParseFiles("template/filter/warning.tpl")
	t.Execute(w, listOut)
}

func funcAbout(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/about.tpl")
	t.Execute(w, nil)
}

func funcTest(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("template/engine/snapshot.tpl")
	t.Execute(w, nil)
}

func config() {
	fmt.Println("Parse config")
	buf, err := ioutil.ReadFile("bearui.conf")
	conf := make(map[string]string)
	if err != nil {
		panic(err)
	} else {
		lines := strings.Split(string(buf), "\x0a")
		for _, line := range lines {
			line = strings.Trim(line, "\x0d\x0a \t")
			if len(line) == 0 || line[0] == '#' {
				continue
			}
			ind := strings.Index(line, "=")
			if ind >= 0 {
				left := strings.Trim(line[:ind], "\x0d\x0a \t")
				right := strings.Trim(line[ind+1:], "\x0d\x0a \t")
				conf[left] = right
			}
		}
	}
	hostELK = conf["elastic"]
	fmt.Println("> Elastic:", hostELK)

	hostDocker = conf["docker"]
	fmt.Println("> Docker:", hostDocker)

	hostGlanceFarm = conf["glancefarm"]
	fmt.Println("> Glance Farm:", hostGlanceFarm)

	hostGlanceProxy = conf["glanceproxy"]
	fmt.Println("> Glance Proxy:", hostGlanceProxy)

	hostEngine = conf["engine"]
	fmt.Println("> Engine:", hostEngine)

	listenHost = conf["host"]
	listenPort = conf["port"]
}

func main() {
	fmt.Println("HoneyBear Web GUI Server")
	var err error

	config()

	//Get docker
	Docker, err = client.NewClient(hostDocker, "v1.25", nil, map[string]string{})
	if err != nil {
		log.Println("Cannot open docker client: ", err)
	}

	//Get ELK
	elkClient, err = elastic.NewClient(elastic.SetURL(hostELK))
	if err != nil {
		log.Println("Cannot open elk client: ", err)
	}

	//css, js, img
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	//test
	http.HandleFunc("/test", funcTest)

	//bearproxy
	http.HandleFunc("/proxy/stat", funcProxy)
	http.HandleFunc("/proxy/stat/live", funcProxyStat)
	http.HandleFunc("/proxy/uid", funcUID)
	http.HandleFunc("/proxy/uid/history", funcUIDHistory)

	//bearfarm
	http.HandleFunc("/farm/host", funcHost)
	http.HandleFunc("/farm/host/stat", funcHostStat)
	http.HandleFunc("/farm/container", funcContainer)
	http.HandleFunc("/farm/container/control", funcContainerControl)
	http.HandleFunc("/farm/container/stat", funcContainerStat)
	http.HandleFunc("/farm/container/livestat", funcContainerLiveStat)
	http.HandleFunc("/farm/topology", funcTopology)

	//bearfilter
	http.HandleFunc("/filter/warning", funcWarning)
	http.HandleFunc("/filter/tracking", funcTracking)
	http.HandleFunc("/filter/tracking/uid", funcTrackingID)
	http.HandleFunc("/filter/tracking/query", funcTrackingIDQuery)
	http.HandleFunc("/filter/tracking/cmd", funcTrackingIDCmd)

	http.HandleFunc("/log", funcLog)

	//bearengine
	http.HandleFunc("/engine/snapshot", funcSnapshot)
	http.HandleFunc("/engine/snapshot/set", funcSnapshotSet)
	http.HandleFunc("/engine/object", funcObject)

	//bearhelp
	http.HandleFunc("/help/about", funcAbout)

	//dashboard, index, handle 404
	http.HandleFunc("/", funcIndex)

	fmt.Printf("Start listen on %s:%s\n", listenHost, listenPort)
	err = http.ListenAndServe(listenHost+":"+listenPort, nil)
	if err != nil {
		log.Println("ListenAndServe: ", err)
	}
}
