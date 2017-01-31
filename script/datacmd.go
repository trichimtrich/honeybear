package main

import (
	"context"
	"encoding/json"
	"fmt"

	elastic "gopkg.in/olivere/elastic.v5"
)

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

func main() {

	elkClient, err := elastic.NewClient(elastic.SetURL("http://192.168.200.1:9200"))
	if err != nil {
		panic(err)
	}

	fmt.Println("[+] ELK Client: OK")
	uid := "a5cb9be512611186"
	fmt.Println("[+] Test UID:", uid)

	reqELK, err := elkClient.Search().Index("request").Type("check").Query(elastic.NewMatchQuery("uid", uid)).Sort("time", false).Size(10000).Do(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("[+] Number request checked:", reqELK.TotalHits())

	numNormal, numQuery := 0, 0
	numTrigger := 0
	for _, item := range reqELK.Hits.Hits {
		var war RequestCheck
		err = json.Unmarshal(*item.Source, &war)
		if err != nil {
			fmt.Println("[Warning] Cannot decode json id", item.Id)
		} else {
			if len(war.Cmd) > 0 {
				numQuery++
				st := war.Cmd[0]    //echo 'bbbb' > comment/'INJECT_HERE'
				st = st[23:]        //echo 'bbbb' > comment/'
				st = st[:len(st)-1] //'
				//=> st = INJECT_HERE
				//st break duoc string sequence => command injection

				b := true
				slash := false
				for i := 0; i < len(st); i++ {
					if slash == true {
						slash = false
						continue
					}
					if st[i] == '\\' {
						slash = true
					}
					if st[i] == '\'' {
						b = false
						break
					}
				}
				if b == false {
					numTrigger++
				}
			} else {
				numNormal++
			}
		}
	}

	fmt.Println("[+] Normal request:", numNormal)
	fmt.Println("[+] Commix request:", numQuery)
	fmt.Println("[+] Commix trigger:", numTrigger)
	fmt.Println("[+] Commix cannot trigger:", numQuery-numTrigger)
}
