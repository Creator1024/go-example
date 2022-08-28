package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadline := time.Now().Add(timeout)

	for tries :=0; time.Now().Before(deadline); tries++{
		_, err := http.Head(url)
		if err == nil {return nil}
		log.Printf("server not responding (%s); retrying...", err)
		time.Sleep(time.Second << uint(tries) ) // 指数退避策略1, 2, 4, 8, 16, 32 ...
	}
	return fmt.Errorf("server %s failed to respond after %s", url, timeout)
}


func main() {
	if err := WaitForServer("http://127.0.0.1:8000"); err != nil{
		log.Fatalf("Site is down: %v\n", err)
	}
}