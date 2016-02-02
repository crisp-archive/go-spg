package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var path *string

func live_handler(w http.ResponseWriter, r *http.Request) {
	dat, err := ioutil.ReadFile(*path)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(w, string(dat))

	if err != nil {
		fmt.Fprintf(w, "The error is %s\n", err)
		return
	}
	fmt.Println("Request", *path)
}

func live_server(dst *string) {
	path = dst
	fmt.Println("Live Mode - Server Init")
	http.HandleFunc("/", live_handler)
	http.ListenAndServe(":8060", nil)
}
