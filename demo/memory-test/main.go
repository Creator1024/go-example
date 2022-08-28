package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	ticker, err := strconv.ParseInt(os.Getenv("Ticker"), 10, 64)
	if err != nil {
		fmt.Println("ENV Ticker not set, use default value(5).")
		ticker = 5
	}
	memoryTicker := time.NewTicker(time.Millisecond * time.Duration(ticker))
	leak := make(map[int][]byte)
	i := 0

	go func() {
		for range memoryTicker.C {
			leak[i] = make([]byte, 1024)
			i++
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	fmt.Println("Strat listening on 0:8000...")
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}

}
