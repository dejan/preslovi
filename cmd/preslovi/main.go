package main

import (
        "fmt"
        "bufio"
        "os"
        "github.com/dejan/preslovi"
)

func main() {
        scanner := bufio.NewScanner(os.Stdin)
        for scanner.Scan() {
            fmt.Println(preslovi.Latinicom(scanner.Text()))
        }
}
