package main

import (
    "flag"
    "fmt"
    "os"

    "github.com/crispgm/go-spg/generator"
    "github.com/crispgm/go-spg/page"
    "github.com/crispgm/go-spg/variables"
)

const VERSION string = "1.0"

func main() {
    currentPath := os.Args[0]
    helpPtr := flag.Bool("help", false, "Help")
    versionPtr := flag.Bool("version", false, "Version Info")
    liveModePtr := flag.Bool("live", false, "Run in live-mode")
    generateModePtr := flag.Bool("generate", false, "Generate File")
    srcPathPtr := flag.String("src", currentPath, "Destination Path")
    dstPathPtr := flag.String("dst", currentPath, "Source Path")

    flag.Usage = usage
    flag.Parse()

    if *versionPtr == true {
        version()
    } else if *helpPtr == true {
        help()
    } else if *liveModePtr == true {
        live(srcPathPtr)
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

func help() {
    fmt.Println("Help")
}

func usage() {
    fmt.Println("SPG Usage:")
    fmt.Println("\t-help")
    fmt.Println("\t-version")
    fmt.Println("\t-compile [-src=SRC_PATH -dst=DST_PATH]")
    fmt.Println("\t-gtype [static|markdown|mdx]")
    fmt.Println("\t-live")
}

func live(srcPathPtr *string) {
    fmt.Println(*srcPathPtr)
}

func generate(srcPathPtr *string, dstPathPtr *string) {
    var jsonBlob = []byte(`{"Value": "Monotremata"}`)
    vrb := variables.New(jsonBlob)

    err, page := page.New("/home/users/zhangwanlong/project/go-spg/testing/favorite-things.md", "/home/users/zhangwanlong/project/go-spg/testing/favorite-things.html", vrb, generator.G_MDX)
    if err != nil {
        panic("Test NewMdx failed")
    }
    page.LoadTemplate()
    err, _ = page.Generate()
    if err != nil {
        panic("Test NewMdx Generate failed")
    }
    page.Output()
}
