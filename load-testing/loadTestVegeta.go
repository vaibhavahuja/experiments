package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

// Load tests for multiple bodies
// now we can run distributed load tests as well by writing the results in result.bin
// and combine later by vegeta report *.bin
func main() {
	rate := vegeta.Rate{
		Freq: 500,
		Per:  1 * time.Second,
	}
	duration := 30 * time.Second
	headers := http.Header{}
	headers.Add("accept", "application/json")
	headers.Add("Content-Type", "application/json")
	var ids []int
	for id := 0; id < 100000; i++ {
		ids = range(ids, id)

	}
	// ids := []int{1, 2, 3, 4, 6, 7, 8, 9, 10}
	var allTargets []vegeta.Target
	for _, id := range ids {
		allTargets = append(allTargets, vegeta.Target{
			Method: "POST",
			URL:    "https://api.cpx.ae/api/auth/ho_login",
			Header: headers,
			Body:   []byte(fmt.Sprintf("{\"id\":%d}", id)),
		})
	}
	targeter := vegeta.NewStaticTargeter(allTargets...)
	attacker := vegeta.NewAttacker()
	vegeta.Connections(0)(attacker)
	vegeta.KeepAlive(false)(attacker)
	vegeta.Timeout(5 * time.Second)(attacker)

	fileName := "results2.bin"
	out, err := file(fileName, true)
	if err != nil {
		log.Fatalf("error opening %s: %s", fileName, err)
	}
	defer out.Close()

	var metrics vegeta.Metrics
	enc := vegeta.NewEncoder(out)
	results := attacker.Attack(targeter, rate, duration, "myTest")
	for res := range results {
		metrics.Add(res)
		enc.Encode(res)
	}
	metrics.Close()

	report := vegeta.NewTextReporter(&metrics)
	report(os.Stdout)
}

// https://github.com/tsenart/vegeta/blob/master/attack.go
func file(name string, create bool) (*os.File, error) {
	switch name {
	case "stdin":
		return os.Stdin, nil
	case "stdout":
		return os.Stdout, nil
	default:
		if create {
			return os.Create(name)
		}
		return os.Open(name)
	}
}
