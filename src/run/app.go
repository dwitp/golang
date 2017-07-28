package main

import (
    "fmt"
    "os"
    "os/exec"
    "log"
    "path"
    "bytes"
)

func main() {
    dirPath, err := os.Getwd()
    if err != nil {
        log.Fatal(err)
    }
    key := path.Base(dirPath)

    cmd := exec.Command("go", "install")
    var out bytes.Buffer
    var stderr bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = &stderr
    err = cmd.Run()
    if err != nil {
        fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
        return
    }
    fmt.Println(out.String())


    if key == "run" {
        log.Fatal("self install - exiting ...")
    }

    _out, err := exec.Command(key).Output()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%s\n", _out)
}