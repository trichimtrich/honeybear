// BearProxy
// Mod from teeproxy
// trichimtrich - phamminhsang
// 1.1
// 24/12/2016

package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/http/httputil"
	"runtime"
	"strings"
	"time"

	elastic "gopkg.in/olivere/elastic.v5"
)

var hostELK = "http://192.168.200.1:9200"
var hostEngine = "http://192.168.200.1:6969"
var ipEngine = "192.168.200.1"
var listenHost = "0.0.0.0"
var listenPort = "8080"
var product = "localhost:8080"
var farm = "localhost:8081"
var timeout1 = 3 //product
var timeout2 = 1 //farm
var tlsPrivKey = ""
var tlsCert = ""

//Error Handler
type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }

var err error

//ELK Service
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

// Handler for httpServer
type handler struct {
	productSV string
	farmSV    string
}

var netClient = &http.Client{
	Timeout: time.Second * 10,
}

// User Identity Container
var IDBucket map[string]map[string]string
var engineID []string
var engineTime int64
var engineDo bool
var engineLearn = true

func GenUID() string {
	u := make([]byte, 8)
	rand.Read(u)
	return hex.EncodeToString(u)
}

func DumpTable() {
	fmt.Printf("%20s %20s %30s %30s %25s %10s %25s\n", "UID", "IP", "Product Session", "Farm Session", "UserAgent", "Status", "Time")
	for k, v := range IDBucket {
		fmt.Printf("%20s %20s %30s %30s %25.20s %10s %25s\n", k, v["ip"], v["product"], v["farm"], v["agent"], v["status"], v["time"])
	}
}

func DupRequest(request *http.Request) (request1 *http.Request, request2 *http.Request) {
	b1 := new(bytes.Buffer)
	b2 := new(bytes.Buffer)
	b3 := new(bytes.Buffer)
	w := io.MultiWriter(b1, b2, b3)
	io.Copy(w, request.Body)
	defer request.Body.Close()
	request.Body = nopCloser{b3}
	dupHeader := make(map[string][]string)
	for k, v := range request.Header {
		dupHeader[k] = v
	}
	request1 = &http.Request{
		Method:        request.Method,
		URL:           request.URL,
		Proto:         request.Proto,
		ProtoMajor:    request.ProtoMajor,
		ProtoMinor:    request.ProtoMinor,
		Header:        request.Header,
		Body:          nopCloser{b1},
		Host:          request.Host,
		ContentLength: request.ContentLength,
		Close:         true,
	}
	request2 = &http.Request{
		Method:        request.Method,
		URL:           request.URL,
		Proto:         request.Proto,
		ProtoMajor:    request.ProtoMajor,
		ProtoMinor:    request.ProtoMinor,
		Header:        dupHeader,
		Body:          nopCloser{b2},
		Host:          request.Host,
		ContentLength: request.ContentLength,
		Close:         true,
	}
	return
}

func BearFilter(w http.ResponseWriter, req *http.Request) bool {
	if strings.Contains(req.RemoteAddr, ipEngine) && req.URL.Path == "/filter" {
		req.ParseForm()
		uid := req.FormValue("uid")
		status := req.FormValue("status")
		if rec, ok := IDBucket[uid]; ok {
			if rec["status"] != status && (rec["status"] < status || rec["status"] == "3") {
				rec["status"] = status

				_, err = elkClient.Index().Index("uid").Type("bear").Id(uid).BodyJson(map[string]string{
					"ip":      rec["ip"],
					"product": rec["product"],
					"farm":    rec["farm"],
					"agent":   rec["agent"],
					"status":  rec["status"],
					"time":    rec["time"],
				}).Refresh("true").Do(context.Background())
				if err != nil {
					fmt.Println("[Filter] Cannot update ELK:", err)
				}
				DumpTable()
				fmt.Println()
			}
		}
		return true
	}
	return false
}

// ServeHTTP - duplicates request / identifies user / controls flow
func (h handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	//BearFilter is here
	if BearFilter(w, req) {
		return
	}

	var productReq, farmReq *http.Request

	timeRequest := time.Now().UTC().Format(time.RFC3339)
	fmt.Println("->", "["+timeRequest+"]", req.RemoteAddr, "http://"+req.Host+req.URL.RequestURI())

	uid := ""
	cookUID, _ := req.Cookie("BearNonce")
	if cookUID != nil {
		uid = cookUID.Value
	}

	session := ""
	cookSession, _ := req.Cookie("PHPSESSID")
	if cookSession != nil {
		session = cookSession.Value
	}

	useragent := req.UserAgent()
	ip := req.RemoteAddr
	ip = ip[:strings.LastIndex(ip, ":")]

	fmt.Println("uid =", uid, "session =", session)

	resetUID := false
	resetSession := false

	savingUser := false

	//Uid co
	if rec, ok := IDBucket[uid]; ok {
		//lay record
		//session khac product ( co the ca 2 cung dang la null )
		if session != rec["product"] {
			//check cookie co input session hay ko, va co trong table ko?
			b := true
			if session != "" {
				for xxx, rec2 := range IDBucket {
					if rec2["product"] == session {
						//Truong hop co, user voi session da bi hijack, mark status =3 , tao record cho hacker
						rec2["status"] = "3"
						//Update vao ELK
						_, err = elkClient.Index().Index("uid").Type("bear").Id(xxx).BodyJson(map[string]string{
							"ip":      rec2["ip"],
							"product": rec2["product"],
							"farm":    rec2["farm"],
							"agent":   rec2["agent"],
							"status":  rec2["status"],
							"time":    rec2["time"],
						}).Refresh("true").Do(context.Background())

						if err != nil {
							fmt.Println("[ELK] Failed to update victim ", xxx, err)
						}

						uid = GenUID()
						IDBucket[uid] = map[string]string{
							"ip":      ip,
							"product": rec2["product"],
							"farm":    rec2["farm"],
							"agent":   useragent,
							"status":  "2",
							"time":    timeRequest}
						b = false
						break
					}
				}
			}

			if b {
				//Truong hop ko co, nhung user dang change cookie, co xu huong hack
				//tao uid moi, reset session ve lai ""
				uid = GenUID()
				IDBucket[uid] = map[string]string{
					"ip":      ip,
					"product": "",
					"farm":    "",
					"agent":   useragent,
					"status":  "1",
					"time":    timeRequest}
				resetSession = true
			}

			resetUID = true

		} else {
			//session giong nhau => check status = 3

			if rec["status"] == "3" { //normal user nay bi mat session, clear cookie
				savingUser = true
			} else { //neu ko phai, tiep tuc check
				//check scheck useragent

				//neu khac useragent, dang bi hack? ko the phan biet User / Hacker
				//mark status 2 = de gui vao farmSV
				if rec["agent"] != useragent {
					IDBucket[uid]["status"] = "2"
					IDBucket[uid]["agent"] = useragent
				}

				//Useragent giong nhau, neu khac ip thi change ip, con ko thi binh thuong
				if rec["ip"] != ip {
					rec["ip"] = ip
				}

			}
		}

	} else { //Uid input ko co trong table

		//php session co
		b := true
		if session != "" {
			for xxx, rec := range IDBucket {
				if rec["product"] == session {
					//Truong hop co, user voi session da bi hijack, mark status =3 , tao record cho hacker
					rec["status"] = "3"
					//Update vao ELK
					_, err = elkClient.Index().Index("uid").Type("bear").Id(xxx).BodyJson(map[string]string{
						"ip":      rec["ip"],
						"product": rec["product"],
						"farm":    rec["farm"],
						"agent":   rec["agent"],
						"status":  rec["status"],
						"time":    rec["time"],
					}).Refresh("true").Do(context.Background())

					if err != nil {
						fmt.Println("[ELK] Failed to update victim ", xxx, err)
					}

					uid = GenUID()
					IDBucket[uid] = map[string]string{
						"ip":      ip,
						"product": rec["product"],
						"farm":    rec["farm"],
						"agent":   useragent,
						"status":  "2",
						"time":    timeRequest}

					b = false
					break
				}
			}
		}

		//session ko co -> tao uid moi
		if b {
			uid = GenUID()
			IDBucket[uid] = map[string]string{
				"ip":      ip,
				"product": "",
				"farm":    "",
				"agent":   useragent,
				"status":  "0",
				"time":    timeRequest}
			resetSession = true
		}

		//ko co uuid thi reset cho no
		resetUID = true
	}

	//session control
	if savingUser { //reset cookie cho user
		w.Header()["Connection"] = []string{"close"}
		w.Header()["Set-Cookie"] = []string{"BearNonce=; Max-Age=0", "PHPSESSID=; Max-Age=0"}
		w.WriteHeader(200)
		w.Write([]byte("<b>Warning:</b> Please change your password, your session has been hacked !"))

	} else {
		//change cookie if need
		cooks := req.Cookies()
		req.Header.Del("Cookie")
		req.Cookie("")
		for _, cook := range cooks {
			if cook.Name == "PHPSESSID" || cook.Name == "BearNonce" {
				continue
			}
			req.AddCookie(cook)
		}

		// Duplicate request
		productReq, farmReq = DupRequest(req)

		//add request to ELK
		reqID := GenUID()
		req.ParseForm()
		_, err = elkClient.Index().Index("request").Type("bear").Id(reqID).BodyJson(map[string]interface{}{
			"uid":    uid,
			"url":    req.URL.Path,
			"agent":  useragent,
			"cookie": cooks,
			"get":    req.URL.Query(),
			"post":   req.PostForm,
			"time":   timeRequest,
		}).Refresh("true").Do(context.Background())

		if err != nil {
			fmt.Println("[ELK] Failed to add new request", reqID, err)
		}

		chanFarm := make(chan bool)

		// Send request to farmSV
		go func() {
			//Recover function
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in f", r)
				}
			}()

			//Add debug header for Scoring engine
			farmReq.Header.Set("RID", reqID)
			if IDBucket[uid]["farm"] != "" {
				farmReq.AddCookie(&http.Cookie{
					Name:  "PHPSESSID",
					Value: IDBucket[uid]["farm"],
				})
			}
			fmt.Println("farm", farmReq.Header["Cookie"])

			// Open new TCP connection to the server
			clientTcpConn, err := net.DialTimeout("tcp", h.farmSV, time.Duration(time.Duration(timeout2)*time.Second))
			if err != nil {
				fmt.Println("[Farm] Failed to connect:", err)
				return
			}
			clientHttpConn := httputil.NewClientConn(clientTcpConn, nil) // Start a new HTTP connection on it
			defer clientHttpConn.Close()                                 // Close the connection to the server
			err = clientHttpConn.Write(farmReq)                          // Pass on the request
			if err != nil {
				fmt.Println("[Farm] Failed to send request:", err)
				return
			}
			farmResp, err := clientHttpConn.Read(farmReq) // Read back the reply
			defer farmResp.Body.Close()

			farmSession := ""
			for _, cook := range farmResp.Cookies() {
				if cook.Name == "PHPSESSID" {
					farmSession = cook.Value
				}
			}
			fmt.Println("Response farm session =", farmSession)

			if farmSession != "" {
				if farmSession == "deleted" {
					IDBucket[uid]["farm"] = ""
				} else {
					IDBucket[uid]["farm"] = farmSession
				}
			}

			//Neu status = 2 thi chi lay response cua FarmSV tra ve
			if IDBucket[uid]["status"] == "2" {
				if resetUID {
					farmResp.Header.Add("Set-Cookie", fmt.Sprintf("BearNonce=%s; Max-Age=31536000", uid))
				}
				if farmSession != "" {
					farmResp.Header.Add("Set-Cookie", fmt.Sprintf("PHPSESSID=%s; Max-Age=31536000", IDBucket[uid]["product"]))
				}
				for k, v := range farmResp.Header {
					w.Header()[k] = v
				}

				w.WriteHeader(farmResp.StatusCode)
				body, _ := ioutil.ReadAll(farmResp.Body)
				w.Write(body)
			}

			chanFarm <- true
			close(chanFarm)
		}()
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in f", r)
			}
		}()

		if IDBucket[uid]["status"] != "2" {
			//Send to productive
			//Append phpsessid if needed
			if IDBucket[uid]["product"] != "" {
				productReq.AddCookie(&http.Cookie{
					Name:  "PHPSESSID",
					Value: IDBucket[uid]["product"],
				})
			}
			fmt.Println("product", productReq.Header["Cookie"])

			// Open new TCP connection to the server
			clientTcpConn, err := net.DialTimeout("tcp", h.productSV, time.Duration(time.Duration(timeout1)*time.Second))
			if err != nil {
				fmt.Println("[Product] Failed to connect:", err)
				return
			}
			clientHttpConn := httputil.NewClientConn(clientTcpConn, nil) // Start a new HTTP connection on it
			defer clientHttpConn.Close()                                 // Close the connection to the server
			err = clientHttpConn.Write(productReq)                       // Pass on the request
			if err != nil {
				fmt.Println("[Product] Failed to send request:", err)
				return
			}
			productResp, err := clientHttpConn.Read(productReq) // Read back the reply
			defer productResp.Body.Close()

			//Check cookie output
			session = ""
			for _, cook := range productResp.Cookies() {
				if cook.Name == "PHPSESSID" {
					session = cook.Value
				}
			}
			fmt.Println("Response product session =", session)

			if session != "" {
				if session == "deleted" {
					IDBucket[uid]["product"] = ""
				} else {
					IDBucket[uid]["product"] = session
				}
			} else {
				if resetSession {
					productResp.Header.Add("Set-Cookie", "PHPSESSID=; Max-Age=0")
				}
			}

			if resetUID {
				productResp.Header.Add("Set-Cookie", fmt.Sprintf("BearNonce=%s; Max-Age=31536000", uid))
			}
			for k, v := range productResp.Header {
				w.Header()[k] = v
			}

			w.WriteHeader(productResp.StatusCode)
			body, _ := ioutil.ReadAll(productResp.Body)
			w.Write(body)

		}

		<-chanFarm
		if len(IDBucket[uid]["product"])*len(IDBucket[uid]["farm"]) == 0 {
			fmt.Println("[FATAL] UID", uid, "sessions are broken")
		}

		//update to ELK
		_, err = elkClient.Index().Index("uid").Type("bear").Id(uid).BodyJson(map[string]string{
			"ip":      IDBucket[uid]["ip"],
			"product": IDBucket[uid]["product"],
			"farm":    IDBucket[uid]["farm"],
			"agent":   IDBucket[uid]["agent"],
			"status":  IDBucket[uid]["status"],
			"time":    timeRequest,
		}).Refresh("true").Do(context.Background())

		if err != nil {
			fmt.Println("[ELK] Failed to update UID:", err)
		}

		//update to ELK History
		_, err = elkClient.Index().Index("uid").Type("history").BodyJson(map[string]string{
			"uid":     uid,
			"ip":      IDBucket[uid]["ip"],
			"product": IDBucket[uid]["product"],
			"farm":    IDBucket[uid]["farm"],
			"agent":   IDBucket[uid]["agent"],
			"status":  IDBucket[uid]["status"],
			"time":    timeRequest,
		}).Refresh("true").Do(context.Background())

		if err != nil {
			fmt.Println("[ELK] Failed to update UID History:", err)
		}

		//Real-time mode
		if !engineLearn {
			//Do engine
			for engineDo {
			}
			engineTime = time.Now().Unix()
			engineID = append(engineID, reqID)
		}

	}
	DumpTable()
	fmt.Println()
}

func LoadRecord() {
	listRec, err := elkClient.Search().Index("uid").Type("bear").Size(10000).Do(context.Background())
	if err != nil {
		log.Println("[ELK] Cannot load UIDs from ELK")
	} else {
		for _, item := range listRec.Hits.Hits {
			var uid UserID
			err = json.Unmarshal(*item.Source, &uid)
			if err != nil {
				log.Println("[ELK] Fail UID", item.Id)
			} else {
				IDBucket[item.Id] = map[string]string{
					"ip":      uid.IP,
					"product": uid.ProductSess,
					"farm":    uid.FarmSess,
					"agent":   uid.UserAgent,
					"status":  uid.Status,
					"time":    uid.Time}
			}
		}
	}
	DumpTable()
	fmt.Println()
}

func DoEngine() {
	for {
		if (len(engineID) >= 10) || ((time.Now().Unix()-engineTime >= 2) && (len(engineID) > 0)) { //2s
			engineDo = true
			par := strings.Join(engineID, ",")
			engineID = engineID[0:0]
			engineDo = false
			fmt.Println(par)
			nresp, err := netClient.Get(hostEngine + "/check?rid=" + par)
			if err != nil {
				fmt.Println("==== ENGINE ==== Failed\n")
				continue
			}
			defer nresp.Body.Close()
			nbody, _ := ioutil.ReadAll(nresp.Body)
			fmt.Println("==== ENGINE ====\n", string(nbody))
		}
	}
}

func config() {
	fmt.Println("Parse config")
	buf, err := ioutil.ReadFile("bearproxy.conf")
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

	hostEngine = conf["engine"]
	fmt.Println("> Engine:", hostEngine)

	ipEngine = conf["engineip"]
	fmt.Println("> Engine IP:", ipEngine)

	product = conf["product"]
	fmt.Println("> Product Server:", product)

	farm = conf["farm"]
	fmt.Println("> Farm Server:", farm)

	listenHost = conf["host"]
	listenPort = conf["port"]

	if conf["mode"] != "learn" {
		engineLearn = false
	}
	fmt.Println("> Learning mode:", engineLearn)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	IDBucket = make(map[string]map[string]string)

	fmt.Println("Bear Proxy!")

	config()
	runtime.GOMAXPROCS(runtime.NumCPU())

	//Get ELK
	elkClient, err = elastic.NewClient(elastic.SetURL(hostELK))
	if err != nil {
		log.Fatal("Cannot open ELK Client: ", err)
	}
	LoadRecord()

	var listener net.Listener

	if len(tlsPrivKey) > 0 {
		cer, err := tls.LoadX509KeyPair(tlsCert, tlsPrivKey)
		if err != nil {
			fmt.Printf("Failed to load certficate: %s and private key: %s\n", tlsCert, tlsPrivKey)
			return
		}

		config := &tls.Config{Certificates: []tls.Certificate{cer}}
		listener, err = tls.Listen("tcp", listenHost+":"+listenPort, config)
		if err != nil {
			fmt.Printf("Failed to listen to %s: %s\n", listenHost+":"+listenPort, err)
			return
		}
	} else {
		listener, err = net.Listen("tcp", listenHost+":"+listenPort)
		if err != nil {
			fmt.Printf("Failed to listen to %s: %s\n", listenHost+":"+listenPort, err)
			return
		}
	}

	h := handler{
		productSV: product,
		farmSV:    farm,
	}

	//Real-time mode
	if !engineLearn {
		//Do Engine
		engineID = make([]string, 0)
		engineTime = time.Now().Unix()
		engineDo = false
		go DoEngine()

	}

	fmt.Println("Running...")
	http.Serve(listener, h)
}
