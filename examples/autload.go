package main

import (
	"fmt"
	_ "github.com/jpfuentes2/env.go/env"
	"os"
)

func main() {
	for _, v := range os.Environ() {
		fmt.Println(v)
	}
}
