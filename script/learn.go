//BearEngine
//Learning phase 1.0

package main

import (
	"context"
	"fmt"
	"time"

	"encoding/json"

	"sort"

	"reflect"

	elastic "gopkg.in/olivere/elastic.v5"
)

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

func main() {
	learnURL := make(map[string]map[string][]string)

	client, err := elastic.NewClient()
	if err != nil {
		panic(err)
	}

	//LEARNING PHASE

	//Create range query from date -> date
	learnStart, _ := time.Parse(time.RFC3339, "2016-12-25T00:00:00Z")
	learnEnd, _ := time.Parse(time.RFC3339, "2016-12-28T00:00:00Z")
	learnDay := elastic.NewRangeQuery("time").Gte(learnStart).Lte(learnEnd)

	//List all documents from query
	learnELK, err := client.Search().Index("request").Type("bear").Query(learnDay).From(0).Size(10000).Do(context.Background())
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
				sort.Strings(learnURL[url][k])
				j := sort.SearchStrings(learnURL[url][k], v[0])
				if j < len(learnURL[url][k]) && learnURL[url][k][j] == v[0] {
					continue
				}
				learnURL[url][k] = append(learnURL[url][k], v[0])
			}
		}
	}

	fmt.Println("---------- Unique records ----------")
	for k, v := range learnURL {
		fmt.Println(k, v)
	}

	//CLASSIFYING PHASE
	//Create range query from date -> date
	classStart, _ := time.Parse(time.RFC3339, "2016-12-25T00:00:00Z")
	classEnd, _ := time.Parse(time.RFC3339, "2016-12-28T00:00:00Z")
	classDay := elastic.NewRangeQuery("time").Gte(classStart).Lte(classEnd)

	//List all documents from query
	classELK, err := client.Search().Index("request").Type("bear").Query(classDay).From(0).Size(10000).Do(context.Background())
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
						sort.Strings(learnURL[url][k])
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

	fmt.Println("---------- Unique actions ----------")

	//Generate
	uniqueURL := make(map[string][]map[string]string)
	for i, item := range classELK.Hits.Hits {
		var curRequest Request
		err = json.Unmarshal(*item.Source, &curRequest)
		if err != nil {
			fmt.Println("Fail:", i, item.Id)
		} else {
			url := curRequest.URL
			if _, ok := learnURL[url]; !ok {
				uniqueURL[url] = make([]map[string]string, 0)
			}
			//map[string]string
			hihi := make(map[string]string)
			for k, v := range curRequest.GetForm { //map[string][]string
				if _, ok := learnURL[url][k]; ok {
					sort.Strings(learnURL[url][k])
					j := sort.SearchStrings(learnURL[url][k], v[0])
					if j < len(learnURL[url][k]) && learnURL[url][k][j] == v[0] {
						hihi[k] = v[0]
					}
				}
			}
			//check if this form already in table
			b := true
			for _, v := range uniqueURL[url] {
				if reflect.DeepEqual(v, hihi) {
					b = false
					break
				}
			}

			//unique
			if b {
				uniqueURL[url] = append(uniqueURL[url], hihi)
			}

		}
	}

	c := 0

	for url, arr := range uniqueURL {
		for _, ar := range arr {
			st := url + " "
			for m, n := range ar {
				st += m + "=" + n + " "
			}
			fmt.Println(st)
			c += 1
		}
	}

	fmt.Println("---------- Total unique ----------", c)

}
