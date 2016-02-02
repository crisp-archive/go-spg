package main

import (
    "flag"
    "fmt"
    "os"

    "github.com/crispgm/go-spg/page"
    "github.com/crispgm/go-spg/variables"
)

const VERSION string = "1.0"

func main() {
    currentPath := os.Args[0]
    versionPtr := flag.Bool("version", false, "Version Info")
    liveModePtr := flag.Bool("live", false, "Run in live-mode")
    generateModePtr := flag.Bool("generate", false, "Generate File")
    enginePtr := flag.String("engine", "mdx", "Engine")
    srcPathPtr := flag.String("src", currentPath, "Destination Path")
    dstPathPtr := flag.String("dst", currentPath, "Source Path")

    flag.Usage = usage
    flag.Parse()

    if *versionPtr == true {
        version()
    } else if *liveModePtr == true {
        live(srcPathPtr, dstPathPtr, enginePtr)
    } else if *generateModePtr == true {
        generate(srcPathPtr, dstPathPtr, enginePtr)
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
    fmt.Println("\t-var=/path/to/json")
    fmt.Println("\t-engine [static|markdown|mdx]")
    fmt.Println("\t-live")
}

func live(srcPathPtr *string, dstPathPtr *string, enginePtr *string) {
    generate(srcPathPtr, dstPathPtr, enginePtr)
    live_server(dstPathPtr)
}

func generate(srcPathPtr *string, dstPathPtr *string, enginePtr *string) {
    var jsonBlob = []byte(``)
    vrb := variables.New(jsonBlob)

    err, page := page.New(*srcPathPtr, *dstPathPtr, vrb, *enginePtr)
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
