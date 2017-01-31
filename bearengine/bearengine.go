//BearEngine
//Core engine

package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"encoding/json"

	"sort"

	"reflect"

	"strings"

	shlex "github.com/flynn-archive/go-shlex"
	"github.com/xwb1989/sqlparser"
	elastic "gopkg.in/olivere/elastic.v5"
)

var hostELK = "http://192.168.200.1:9200"
var hostFilter = "http://192.168.200.105:80/filter"
var listenHost = "0.0.0.0"
var listenPort = "6969"

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

var netClient = &http.Client{
	Timeout: time.Second * 10,
}
var elkClient *elastic.Client
var err error
var curProfileID string
var curProfile Snapshot
var learnURL map[string]map[string][]string

type QueryToken struct {
	Token   int
	Content string
}

type QueryObject struct {
	Query  string
	Error  string
	Tokens []QueryToken
}

type CmdObject struct {
	Command string
	Tokens  []string
}

type WebObject struct {
	URL       string
	Param     map[string]string
	ReqID     []string
	Queries   [][]QueryObject
	Cmds      [][]CmdObject
	NullQuery bool
	NullCmd   bool
}

var learnObj []WebObject

func LearnMode() {
	learnURL = make(map[string]map[string][]string)

	//LEARNING PHASE

	//Create range query from date -> date
	learnStart, _ := time.Parse(time.RFC3339, curProfile.LearnStart)
	learnEnd, _ := time.Parse(time.RFC3339, curProfile.LearnEnd)
	learnDay := elastic.NewRangeQuery("time").Gte(learnStart).Lte(learnEnd)

	//List all documents from query
	learnELK, err := elkClient.Search().Index("request").Type("bear").Query(learnDay).From(0).Size(10000).Do(context.Background())
	if err != nil {
		panic(err)
	}

	fmt.Println("Total hit:", learnELK.TotalHits(), "- Total time ms:", learnELK.TookInMillis)
	for i, item := range learnELK.Hits.Hits {
		var curRequest Request
		err = json.Unmarshal(*item.Source, &curRequest)
		if err != nil {
			fmt.Println("Fail:", i, item.Id)
		} else {
			url := curRequest.URL
			if _, ok := learnURL[url]; !ok {
				learnURL[url] = make(map[string][]string)
			}
			for k, v := range curRequest.GetForm {
				if _, ok := learnURL[url][k]; !ok {
					learnURL[url][k] = make([]string, 0)
				}
				j := sort.SearchStrings(learnURL[url][k], v[0])
				if j < len(learnURL[url][k]) && learnURL[url][k][j] == v[0] {
					continue
				}
				learnURL[url][k] = append(learnURL[url][k], v[0])
				sort.Strings(learnURL[url][k])
			}
		}
	}

	fmt.Println("---------- Unique records ----------")
	for k, v := range learnURL {
		fmt.Println(k, v)
	}

	//CLASSIFYING PHASE
	//Create range query from date -> date
	classStart, _ := time.Parse(time.RFC3339, curProfile.VerifyStart)
	classEnd, _ := time.Parse(time.RFC3339, curProfile.VerifyEnd)
	classDay := elastic.NewRangeQuery("time").Gte(classStart).Lte(classEnd)

	//List all documents from query
	classELK, err := elkClient.Search().Index("request").Type("bear").Query(classDay).From(0).Size(10000).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("\nTotal hit:", classELK.TotalHits(), "- Total time ms:", classELK.TookInMillis)
	//Filter
	for i, item := range classELK.Hits.Hits {
		var curRequest Request
		err = json.Unmarshal(*item.Source, &curRequest)
		if err != nil {
			fmt.Println("Fail:", i, item.Id)
		} else {
			url := curRequest.URL
			if _, ok := learnURL[url]; !ok {
				fmt.Println("No URL record:", url)
			} else {
				//map[string][]string
				for k, v := range curRequest.GetForm {
					if _, ok := learnURL[url][k]; ok {
						//sort.Strings(learnURL[url][k])
						j := sort.SearchStrings(learnURL[url][k], v[0])
						if j < len(learnURL[url][k]) && learnURL[url][k][j] == v[0] {
							continue
						}
						fmt.Println(">Remove", url, k)
						delete(learnURL[url], k)
					}
				}
			}
		}
	}

	//Generate object
	learnObj = make([]WebObject, 0)
	for _, item := range classELK.Hits.Hits {
		var curRequest Request
		err = json.Unmarshal(*item.Source, &curRequest)
		if err != nil {
			fmt.Println("Fail:", item.Id)
		} else {
			//Check If request URL already learn
			if _, ok := learnURL[curRequest.URL]; !ok {
				continue
			}

			//Create new object
			var webObj WebObject
			webObj.URL = curRequest.URL
			webObj.Param = make(map[string]string)
			for k, v := range curRequest.GetForm {
				if _, ok := learnURL[webObj.URL][k]; ok {
					//sort.Strings(learnURL[webObj.URL][k])
					j := sort.SearchStrings(learnURL[webObj.URL][k], v[0])
					if j < len(learnURL[webObj.URL][k]) && learnURL[webObj.URL][k][j] == v[0] {
						webObj.Param[k] = v[0]
					}
				}
			}

			//Check if object already
			refWebObj := &webObj
			isNew := true
			for l, obj := range learnObj {
				if obj.URL == webObj.URL && reflect.DeepEqual(obj.Param, webObj.Param) {
					isNew = false
					refWebObj = &learnObj[l]
					break
				}
			}

			//If new , init new array
			if isNew {
				webObj.ReqID = make([]string, 0)
				webObj.Queries = make([][]QueryObject, 0)
				webObj.Cmds = make([][]CmdObject, 0)
				webObj.NullQuery = false
				webObj.NullCmd = false
			}

			//Analysis new request

			refWebObj.ReqID = append(refWebObj.ReqID, item.Id) //Add new rid
			if curRequest.URL == "/index.php" {
				//fmt.Println(refWebObj, item.Id, refWebObj.ReqID, isNew)
			}
			//add new queries
			queryELK, err := elkClient.Search().Index("query").Type("bear").Query(elastic.NewMatchQuery("rid", item.Id)).Size(10000).Do(context.Background())
			if err != nil {
				fmt.Println("[Learn] RID", item.Id, "cannot get queries ELK")
			} else {
				qrObjs := make([]QueryObject, 0)
				for _, itemQr := range queryELK.Hits.Hits {
					var qr Query
					err = json.Unmarshal(*itemQr.Source, &qr)
					if err != nil {
						fmt.Println("[Learn] RID", item.Id, "- Query ID", itemQr.Id, "cannot decode json")
					} else {
						var qrObj QueryObject
						ParseSQL(qr.Query, &qrObj)
						qrObjs = append(qrObjs, qrObj)
					}
				}
				if len(qrObjs) > 0 {
					refWebObj.Queries = append(refWebObj.Queries, qrObjs)
				} else {
					refWebObj.NullQuery = true
				}
			}

			//add new commands
			cmdELK, err := elkClient.Search().Index("cmd").Type("bear").Query(elastic.NewMatchQuery("rid", item.Id)).Size(10000).Do(context.Background())
			if err != nil {
				fmt.Println("[Learn] RID", item.Id, "cannot get commands ELK")
			} else {
				cmdObjs := make([]CmdObject, 0)
				for _, itemCmd := range cmdELK.Hits.Hits {
					var cmd Command
					err = json.Unmarshal(*itemCmd.Source, &cmd)
					if err != nil {
						fmt.Println("[Learn] RID", item.Id, "- Command ID", itemCmd.Id, "cannot decode json")
					} else {
						var cmdObj CmdObject
						ParseCMD(cmd.Cmd, &cmdObj)
						cmdObjs = append(cmdObjs, cmdObj)
					}
				}
				if len(cmdObjs) > 0 {
					refWebObj.Cmds = append(refWebObj.Cmds, cmdObjs)
				} else {
					refWebObj.NullCmd = true
				}
			}

			//If new, append to learnObj
			if isNew {
				learnObj = append(learnObj, webObj)
			}

		}
	}

	c := 0
	for _, item := range learnObj {
		fmt.Println(item.URL, item.Param)
		c += len(item.ReqID)
		//fmt.Println(item.Queries)
		//fmt.Println(item.Cmds)
		//fmt.Println()
	}

	fmt.Println("---------- Total unique ----------", len(learnObj), c)

}

func userCheck(uid string, status string) {
	_, err = netClient.Get(hostFilter + "?uid=" + uid + "&status=" + status)
	if err != nil {
		fmt.Println("User", uid, "cannot update status", status)
	} else {
		fmt.Println("User", uid, "update to status", status)
	}
}

func addCheck(t RequestCheck) {
	_, err := elkClient.Index().Index("request").Type("check").BodyJson(t).Refresh("true").Do(context.Background())
	if err != nil {
		fmt.Println("[Check] Cannot add new Request ELK:", err)
	}
}

func funcCheck(w http.ResponseWriter, r *http.Request) {
	//Using record
	//rid := "ad9bf2ae8a5992bc" //hehe.php
	//rid := "21c05dcd5b84b45c" //customer.php?hehe=123&do=add
	//rid := "34ded652e59c63f1" //statement - normal
	//rid := "9ff53052ee7c9a9e" //statement - sqli
	//rid := "b766fedde996a8ae" //command - normal
	//rid := "08f3a8797dfa7d37" //command - inject
	r.ParseForm()
	rids := strings.Split(r.FormValue("rid"), ",")
	for _, rid := range rids {
		fmt.Println("==== New request ====", rid)

		//Create request Check
		reqCheck := RequestCheck{}
		reqCheck.Time = time.Now().UTC().Format(time.RFC3339)
		reqCheck.UID = ""
		reqCheck.IP = ""
		reqCheck.Status = ""
		reqCheck.RID = rid
		reqCheck.URL = ""
		reqCheck.Param = make(map[string]string)
		reqCheck.Violation = ""
		reqCheck.Query = make([]string, 0)
		reqCheck.Cmd = make([]string, 0)
		reqCheck.WebObject = make(map[string]string)
		reqCheck.QueryObject = make([][]string, 0)
		reqCheck.CmdObject = make([][]string, 0)

		reqELK, err := elkClient.Get().Index("request").Type("bear").Id(rid).Do(context.Background())
		if err != nil {
			fmt.Println("==== End Request ==== RID not found\n") //ko tim thay request
			w.Write([]byte(rid + " - RID not found\n"))

			reqCheck.Violation = "rid"
			addCheck(reqCheck) //request Check
			//no UID for update
			continue
		}

		var req Request
		_ = json.Unmarshal(*reqELK.Source, &req)
		fmt.Println("URL:", req.URL)

		reqCheck.URL = req.URL //request Check
		reqCheck.UID = req.UID //request Check
		uidELK, err := elkClient.Get().Index("uid").Type("bear").Id(req.UID).Do(context.Background())
		if err != nil {
			fmt.Println("[Check] Cannot find UID ELK:", err)
		} else {
			var uid UserID
			_ = json.Unmarshal(*uidELK.Source, &uid)
			reqCheck.IP = uid.IP
			reqCheck.Status = uid.Status
		}

		if _, ok := learnURL[req.URL]; !ok {
			fmt.Println("==== End Request ==== 404 URL\n") //url khong ton tai
			w.Write([]byte(rid + " - 404 URL\n"))

			reqCheck.Violation = "404"
			addCheck(reqCheck) //request Check
			//404, normal
			continue
		}

		//check web object
		param := make(map[string]string)
		for k, v := range req.GetForm {
			reqCheck.Param[k] = v[0] //request Check
			if _, ok := learnURL[req.URL][k]; ok {
				j := sort.SearchStrings(learnURL[req.URL][k], v[0])
				if j < len(learnURL[req.URL][k]) && learnURL[req.URL][k][j] == v[0] {
					param[k] = v[0]
					reqCheck.WebObject[k] = v[0] //request Check
				}
			}
		}

		//check web action
		isAction := false
		var webObj WebObject
		for _, obj := range learnObj {
			if obj.URL == req.URL && reflect.DeepEqual(obj.Param, param) {
				fmt.Println("Param:", obj.Param) //la web action
				webObj = obj
				isAction = true
				break
			}
		}

		if isAction {
			//check query object
			queryELK, err := elkClient.Search().Index("query").Type("bear").Query(elastic.NewMatchQuery("rid", rid)).Size(10000).Do(context.Background())
			if err != nil {
				fmt.Println("[Check] RID", rid, "cannot get queries ELK")
			} else {
				qrObjs := make([]QueryObject, 0)
				for _, itemQr := range queryELK.Hits.Hits {
					var qr Query
					err = json.Unmarshal(*itemQr.Source, &qr)
					if err != nil {
						fmt.Println("[Check] RID", rid, "- Query ID", itemQr.Id, "cannot decode json")
					} else {
						var qrObj QueryObject
						ParseSQL(qr.Query, &qrObj)
						qrObjs = append(qrObjs, qrObj)
						reqCheck.Query = append(reqCheck.Query, qr.Query) //request Check
					}
				}
				//check null query first
				isQuery := (len(qrObjs) == 0) && webObj.NullQuery
				if len(qrObjs) > 0 {
					fmt.Println(qrObjs)
					for l, checkObj := range webObj.Queries {
						reqCheck.QueryObject = append(reqCheck.QueryObject, make([]string, 0)) //request Check

						//array
						if len(checkObj) == len(qrObjs) {
							subQuery := true
							fmt.Println(checkObj)
							for i := 0; i < len(checkObj); i++ {
								reqCheck.QueryObject[l] = append(reqCheck.QueryObject[l], checkObj[i].Query) //request Check

								if checkObj[i].Error != qrObjs[i].Error || len(checkObj[i].Tokens) != len(qrObjs[i].Tokens) {
									subQuery = false
									break
								}
								for j := 0; j < len(checkObj[i].Tokens); j++ {
									if checkObj[i].Tokens[j].Token != qrObjs[i].Tokens[j].Token {
										subQuery = false
										break
									}
								}
								if !subQuery {
									break
								}
							}
							if subQuery {
								isQuery = true
								break
							}
						}
					}

				}
				if !isQuery {
					w.Write([]byte(rid + " - Sql Injection\n"))
					fmt.Println("==== End Request ==== Sql Injection\n")

					reqCheck.Violation = "query" //request Check
					addCheck(reqCheck)
					userCheck(reqCheck.UID, "2") //hacker
					continue
				}

			}

			//check cmd object
			cmdELK, err := elkClient.Search().Index("cmd").Type("bear").Query(elastic.NewMatchQuery("rid", rid)).Size(10000).Do(context.Background())
			if err != nil {
				fmt.Println("[Learn] RID", rid, "cannot get commands ELK")
			} else {
				cmdObjs := make([]CmdObject, 0)
				for _, itemCmd := range cmdELK.Hits.Hits {
					var cmd Command
					err = json.Unmarshal(*itemCmd.Source, &cmd)
					if err != nil {
						fmt.Println("[Learn] RID", rid, "- Command ID", itemCmd.Id, "cannot decode json")
					} else {
						var cmdObj CmdObject
						ParseCMD(cmd.Cmd, &cmdObj)
						cmdObjs = append(cmdObjs, cmdObj)
						reqCheck.Cmd = append(reqCheck.Cmd, cmd.Cmd) //request Check
					}
				}
				//check null cmd first
				isCmd := (len(cmdObjs) == 0) && webObj.NullCmd
				if len(cmdObjs) > 0 { //check cmd
					//fmt.Println(cmdObjs)
					for l, checkObj := range webObj.Cmds {
						reqCheck.CmdObject = append(reqCheck.CmdObject, make([]string, 0)) //request Check

						//array
						if len(checkObj) == len(cmdObjs) {
							subCmd := true
							//fmt.Println(checkObj)
							for i := 0; i < len(checkObj); i++ {
								reqCheck.CmdObject[l] = append(reqCheck.CmdObject[l], checkObj[i].Command) //request Check

								if (len(checkObj[i].Tokens) != len(cmdObjs[i].Tokens)) || (checkObj[i].Tokens[len(checkObj[i].Tokens)-1] != cmdObjs[i].Tokens[len(cmdObjs[i].Tokens)-1]) {
									subCmd = false
									break
								}
							}
							if subCmd {
								isCmd = true
								break
							}
						}
					}
				} //check null cmd
				if !isCmd {
					w.Write([]byte(rid + " - Command Injection\n"))
					fmt.Println("==== End Request ==== Command Injection\n")

					reqCheck.Violation = "cmd"
					addCheck(reqCheck)           //request Check
					userCheck(reqCheck.UID, "2") //hacker
					continue
				}
			}

			w.Write([]byte(rid + " - Good Request\n"))
			fmt.Println("==== End Request ==== Good Request\n")

			reqCheck.Violation = "good"
			addCheck(reqCheck) //request Check
			//normal
		} else {
			w.Write([]byte(rid + " - Web Object Not Found\n"))
			fmt.Println("==== End Request ==== Web Object Not Found\n")

			reqCheck.Violation = "web"
			addCheck(reqCheck)           //request Check
			userCheck(reqCheck.UID, "1") //suspect
		}
	}
}

func funcGet(w http.ResponseWriter, r *http.Request) {
	//Get records
	b, _ := json.Marshal(learnObj)
	w.Header()["Content-type"] = []string{"application/json"}
	w.Write(b)
}

func funcSet(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method == "GET" && r.FormValue("id") != "" {
		id := r.FormValue("id")
		proELK, err := elkClient.Get().Index("snapshot").Type("bear").Id(id).Do(context.Background())
		if err != nil {
			fmt.Println("[Set Snapshot] Failed to get ELK:", err)
		} else {
			var profile Snapshot
			err = json.Unmarshal(*proELK.Source, &profile)
			if err != nil {
				fmt.Println("[Set Snapshot] Failed to decode json:", err)
			} else {
				curProfileID = id
				curProfile = profile
				fmt.Println()
				fmt.Println("Current profile:", curProfileID)
				fmt.Println("Name:", curProfile.Name)
				fmt.Println("Desc:", curProfile.Description)
				fmt.Println("Learn:", curProfile.LearnStart, curProfile.LearnEnd)
				fmt.Println("Verify:", curProfile.VerifyStart, curProfile.VerifyEnd)
				LearnMode()
				w.Write([]byte("ok"))
				return
			}
		}
	}
	w.Write([]byte("fail"))
}

func funcCurrent(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(curProfileID))
}

func funcIndex(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hihi"))
}

func ParseSQL(query string, qrObj *QueryObject) {
	//select *,123,456,'7a8xxxx/*aahihi\'*/xxxx9'/*hihi*/,"zzzz", ` + "`aaaa`" + ` from information_schema.tables, hihi, (select * from master_db) where 1=2 or hihi>(select 100 from abcd) union select * from information_schema.columns where (select count(*) from abcd)>100
	_, err := sqlparser.Parse(query)
	if err != nil {
		fmt.Println("[SQL Parser]", err)
	}
	qrObj.Error = fmt.Sprintf("%s", err)
	qrObj.Query = query
	qrObj.Tokens = make([]QueryToken, 0)
	tokens := sqlparser.NewStringTokenizer(query)
	for {
		token, content := tokens.Scan()
		if token == 0 {
			break
		}
		qrObj.Tokens = append(qrObj.Tokens, QueryToken{Token: token, Content: string(content)})
		//fmt.Println(token, string(content))
	}
}

func ParseCMD(cmd string, cmdObj *CmdObject) {
	//echo 'bbb' > comment/'aaaa'
	cmdObj.Tokens = make([]string, 0)
	cmdObj.Command = cmd

	ioBuf := new(bytes.Buffer)
	ioBuf.Write([]byte(cmd))

	tokens, err := shlex.NewLexer(ioBuf)
	if err != nil {
		fmt.Println("Bi loi cmd roi", err)
	} else {
		for {
			token, err := tokens.NextWord()
			if err != nil {
				cmdObj.Tokens = append(cmdObj.Tokens, fmt.Sprintf("%s", err))
				break
			} else {
				cmdObj.Tokens = append(cmdObj.Tokens, token)
			}
		}
	}
}

func config() {
	fmt.Println("Parse config")
	buf, err := ioutil.ReadFile("bearengine.conf")
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

	hostFilter = conf["filter"]
	fmt.Println("> Filter:", hostFilter)

	listenHost = conf["host"]
	listenPort = conf["port"]
}

func main() {
	config()

	elkClient, err = elastic.NewClient(elastic.SetURL(hostELK))
	if err != nil {
		panic(err)
	}

	//get first profile
	proELK, err := elkClient.Search().Index("snapshot").Type("bear").Sort("time", false).Size(1).Do(context.Background())
	if err != nil {
		fmt.Println(err)
	} else {
		if proELK.TotalHits() == 0 {
			fmt.Println("No snapshot found")
		} else {
			err = json.Unmarshal(*proELK.Hits.Hits[0].Source, &curProfile)
			if err != nil {
				panic("Cannot decode json")
			}
			curProfileID = proELK.Hits.Hits[0].Id
			fmt.Println("Current profile:", curProfileID)
			fmt.Println("Name:", curProfile.Name)
			fmt.Println("Desc:", curProfile.Description)
			fmt.Println("Learn:", curProfile.LearnStart, curProfile.LearnEnd)
			fmt.Println("Verify:", curProfile.VerifyStart, curProfile.VerifyEnd)
			LearnMode()
		}
	}

	//test parser
	//ParseCMD("echo 123 > comment/hoho")

	//web api
	http.HandleFunc("/check", funcCheck)
	http.HandleFunc("/get", funcGet)
	http.HandleFunc("/set", funcSet)
	http.HandleFunc("/current", funcCurrent)
	http.HandleFunc("/", funcIndex)

	fmt.Printf("Start listen on %s:%s\n", listenHost, listenPort)
	err = http.ListenAndServe(listenHost+":"+listenPort, nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}
