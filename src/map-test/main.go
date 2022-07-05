package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title string
	Year  int  `json:"released"`
	Color bool `json:"color,omitempty"`
}

func main() {

	movies := []Movie{
		{"title1", 2000, true},
		{"title2", 2001, false},
	}

	data, err := json.Marshal(movies)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", data)
	fmt.Printf("%T\n", data)
	// [{"Title":"title1","released":2000,"color":true},{"Title":"title2","released":2001}]

	dataIndent, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", dataIndent)
	//[
	//	{
	//	"Title": "title1",
	//	"released": 2000,
	//	"color": true
	//	},
	//	{
	//	"Title": "title2",
	//	"released": 2001
	//	}
	//]

	var m []Movie
	if err := json.Unmarshal(data, &m); err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%v", m)
	// [{title1 2000 true} {title2 2001 false}]
}
