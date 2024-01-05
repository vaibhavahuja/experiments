package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

// Load tests for multiple bodies
func main() {
	rate := vegeta.Rate{
		Freq: 10000,
		Per:  1 * time.Second,
	}
	duration := 5 * time.Second
	headers := http.Header{}
	headers.Add("accept", "application/json")
	headers.Add("Content-Type", "application/json")

	ids := []int{1, 2, 3, 4, 6, 7, 8, 9, 10}
	var allTargets []vegeta.Target
	for _, id := range ids {
		allTargets = append(allTargets, vegeta.Target{
			Method: "POST",
			URL:    "http://localhost:8080/hello",
			Header: headers,
			Body:   []byte(fmt.Sprintf("{\"id\":%d}", id)),
		})
	}
	targeter := vegeta.NewStaticTargeter(allTargets...)
	attacker := vegeta.NewAttacker()
	vegeta.Connections(0)(attacker)
	vegeta.KeepAlive(false)(attacker)
	vegeta.Timeout(5 * time.Second)(attacker)

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "myTest") {
		metrics.Add(res)
	}
	metrics.Close()

	report := vegeta.NewTextReporter(&metrics)
	report(os.Stdout)
}
