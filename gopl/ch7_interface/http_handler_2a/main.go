package main

import (
	"fmt"
	"log"
	"net/http"
)

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db database) get(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	price, ok := db[item]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "item %s not found", item)
		return
	}
	fmt.Fprintf(w, "price of item %s is: %s", item, price)
}

func main() {
	//var err error = syscall.Errno(2)
	//fmt.Println(err)
	//fmt.Println(err.Error())
	db := database{"shoes": 20, "socks": 2}
	httpAddr := "localhost:8000"
	mux := http.NewServeMux()
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/get", db.get)
	fmt.Printf("starting http server, listen on %s ...\n", httpAddr)
	log.Fatal(http.ListenAndServe(httpAddr, mux))
}
