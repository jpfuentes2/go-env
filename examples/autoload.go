package main

import (
	"fmt"
	_ "github.com/jpfuentes2/go-env/autoload"
	"os"
)

func main() {
	for _, v := range os.Environ() {
		fmt.Println(v)
	}
}
