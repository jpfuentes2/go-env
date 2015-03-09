package main

import (
	"fmt"
	"github.com/jpfuentes2/go-env"
	"os"
	"path"
)

// Note this is exactly what autoload.go does
func main() {
	pwd, _ := os.Getwd()
	file := path.Join(pwd, ".env")
	env.ReadEnv(file)
	for _, v := range os.Environ() {
		fmt.Println(v)
	}
}
