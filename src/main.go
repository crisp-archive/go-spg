package main

import (
    "flag"
    "fmt"
    "os"
)

func main() {
    currentPath := os.Args[0]
    configPathPtr := flag.String("config", currentPath, "Config Path")
    outputPathPtr := flag.String("output", currentPath, "Output Path")
    versionPtr := flag.Bool("version", false, "Version Info")

    flag.Parse()

    fmt.Println(currentPath)
    fmt.Println(*configPathPtr)
    fmt.Println(*outputPathPtr)
    fmt.Println(*versionPtr)

    fmt.Println(flag.Args())
}
