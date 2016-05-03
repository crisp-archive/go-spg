package main

import (
	"flag"
	"fmt"
	"os"
	"io/ioutil"
	"net/http"

	"github.com/crispgm/go-spg/page"
)

const VERSION string = "1.0"

var path *string

func main() {
	currentPath := os.Args[0]
	versionPtr := flag.Bool("version", false, "Version Info")
	liveModePtr := flag.Bool("live", false, "Run in live-mode")
	generateModePtr := flag.Bool("generate", false, "Generate File")
	srcPathPtr := flag.String("src", currentPath, "Destination Path")
	dstPathPtr := flag.String("dst", currentPath, "Source Path")

	flag.Usage = usage
	flag.Parse()

	if *versionPtr == true {
		version()
	} else if *liveModePtr == true {
		live(srcPathPtr, dstPathPtr)
	} else if *generateModePtr == true {
		generate(srcPathPtr, dstPathPtr)
	} else {
		usage()
	}
}

func version() {
	fmt.Printf("SPG %s, single page generator.\n", VERSION)
	fmt.Println("Copyright (c) David Zhang, 2015")
}

func usage() {
	fmt.Println("SPG Usage:")
	fmt.Println("\t-version")
	fmt.Println("\t-compile [-src=SRC_PATH -dst=DST_PATH]")
	fmt.Println("\t-live")
}

func live(srcPathPtr *string, dstPathPtr *string) {
	generate(srcPathPtr, dstPathPtr)
	live_server(dstPathPtr)
}

func generate(srcPathPtr *string, dstPathPtr *string) {
	err, page := page.New(*srcPathPtr, *dstPathPtr)
	if err != nil {
		panic("Page Init Failed")
	}
	page.LoadTemplate()
	err, _ = page.Generate()
	if err != nil {
		panic("Page Generate Failed")
	}
	page.Output()
}

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
