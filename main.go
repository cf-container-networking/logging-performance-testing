package main

import (
	"fmt"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

func main() {
	rate := uint64(2000) // per second
	duration := 60 * time.Second
	targeter := vegeta.NewStaticTargeter(vegeta.Target{
		Method: "GET",
		URL:    "http://proxy.slouch.c2c.cf-app.com/proxy/google.com",
	})
	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration) {
		metrics.Add(res)
	}
	metrics.Close()

	fmt.Printf("99th percentile: %s\n", metrics.Latencies.P99)
	fmt.Printf("Mean is: %s\n", metrics.Latencies.Mean)
	fmt.Printf("Bytes In: %+v\n", metrics.BytesIn)
	fmt.Printf("Bytes Out: %+v\n", metrics.BytesOut)
	fmt.Printf("Requests: %+v\n", metrics.Requests)
	fmt.Printf("Rate: %+v\n", metrics.Rate)
	fmt.Printf("Success: %+v\n", metrics.Success)
	fmt.Printf("Completed Requests: %+v\n", metrics.Success*float64(metrics.Requests))
}
