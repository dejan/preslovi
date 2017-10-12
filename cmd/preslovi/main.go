package main

import (
	"bufio"
	"fmt"
	"github.com/dejan/preslovi"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(preslovi.Latinicom(scanner.Text()))
	}
}
