package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			fmt.Fprintf(w, "%s: %s\n", item, price)
		}
	case "/get":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "item %s not found", item)
			return
		}
		fmt.Fprintf(w, "price of item %s is: %s", item, price)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "page %s not found\n", req.URL)
	}
}

func main() {
	db := database{"shoes": 20, "socks": 2}
	httpAddr := "localhost:8000"
	fmt.Printf("starting http server, listen on %s ...\n", httpAddr)
	log.Fatal(http.ListenAndServe(httpAddr, db))
}
