package main

import (
    "fmt"
    "os"
    "os/exec"
    "log"
    "path"
)

func main() {
    dirPath, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }
    key := path.Base(dirPath)

    cmd := exec.Command("go", "install")
    err = cmd.Start()
    if err != nil {
        log.Fatal(err)
    }


    if key == "run" {
        log.Fatal("self install - exiting ...")
    }
    out, err := exec.Command(key).Output()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s\n", out)
}