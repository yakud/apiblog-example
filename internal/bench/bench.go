package bench

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var client = &http.Client{}

func Bench() {
	for {
		t := time.Now()
		queryCreate()
		fmt.Println(time.Now().Sub(t))
	}
}

func queryCreate() {
	url := "http://127.0.0.1:8080"

	var jsonStr = []byte(`{"query": "mutation{create(name: \"hello\", shortDescr: \"desc\"){ id }}"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}
