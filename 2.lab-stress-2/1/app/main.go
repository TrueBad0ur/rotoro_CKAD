package main

import (
    "fmt"
    "math/rand"
    "os/exec"
)

func main() {
    if rand.Intn(2) == 1 {
	    fmt.Println("True!")
        cmd := exec.Command("touch", "/tmp/test")
        stdout, err := cmd.Output()
        _ = stdout

        if err != nil {
            fmt.Println(err.Error())
            return
        }

        cmd1 := exec.Command("sleep", "100000")
        stdout1, err1 := cmd1.Output()
        _ = stdout1

        if err1 != nil {
            fmt.Println(err1.Error())
            return
        }
    } else {
        fmt.Println("False!")
    }
}
