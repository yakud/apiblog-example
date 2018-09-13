package bench

import (
	"bytes"
	"fmt"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/montanaflynn/stats"
)

var client = &http.Client{}

func durationsToFloat64(durations []time.Duration) []float64 {
	data := make([]float64, len(durations))
	for i, j := range durations {
		data[i] = float64(j.Nanoseconds() / int64(time.Millisecond))
	}

	return data
}

func Bench() {
	var count int32
	durations := make([]time.Duration, 0)

	go func() {
		fmt.Println(durations)
		data := durationsToFloat64(durations)
		p50, _ := stats.Percentile(data, 50)
		p75, _ := stats.Percentile(data, 75)
		p95, _ := stats.Percentile(data, 95)
		p99, _ := stats.Percentile(data, 99)

		ticker := time.NewTicker(time.Second)
		for range ticker.C {
			fmt.Printf(
				"%d events/sec; p50: %d\tp75: %d\tp95: %d\tp99: %d\n",
				count,
				int(p50),
				int(p75),
				int(p95),
				int(p99),
			)

			atomic.StoreInt32(&count, 0)
			durations = durations[:0]
		}
	}()

	for {
		t := time.Now()
		queryCreate()

		durations = append(durations, time.Now().Sub(t))
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
