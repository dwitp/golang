package u

import (
    "fmt"
)

func D(args ...interface{}) {
    fmt.Println("-----------------------------------------------------")
    for i := range args {
        fmt.Printf("{%T} ", args[i])
        fmt.Printf("%+v\n", args[i])
    }
    fmt.Println("-----------------------------------------------------")
    fmt.Println()
}


func S(x []int) {
    fmt.Printf("len=%d cap=%d %v\n", len(x), cap(x), x)
}
