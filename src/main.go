package main

import (
    "flag"
    "fmt"
    "os"
)

const VERSION string = "1.0"

func main() {
    currentPath := os.Args[0]
    helpPtr := flag.Bool("help", false, "Help")
    versionPtr := flag.Bool("version", false, "Version Info")
    liveModePtr := flag.Bool("live", false, "Run in live-mode")
    generateModePtr := flag.Bool("generate", false, "Generate Static Site")
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
    fmt.Println("\t-output")
    fmt.Println("\t-var-{name}={value}")
}

func live(srcPathPtr *string) {
    fmt.Println(*srcPathPtr)
}

func generate(srcPathPtr *string, dstPathPtr *string) {
    fmt.Println(*srcPathPtr)
    fmt.Println(*dstPathPtr)
}
