package bench

import (
	"bytes"
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var client = &http.Client{}

func Bench() {
	var count int32

	go func() {
		ticker := time.NewTicker(time.Second)
		for _ := range ticker.C {
			fmt.Println(count, "events/sec")
			atomic.StoreInt32(&count, 0)
		}
	}()

	for {
		t := time.Now()
		queryCreate()
		fmt.Println(time.Now().Sub(t))
		atomic.AddInt32(&count, 1)
	}
}

func queryCreate() {
	url := "http://127.0.0.1:8080/graphql"

	var jsonStr = []byte(`{"query": "mutation{create(name: \"hello\", shortDescr: \"desc\"){ id }}"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
}
