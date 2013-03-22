package main

import (
	"fmt"
	"github.com/jpfuentes2/env.go/env"
	"os"
	"path"
)

// Note this is exactly what autoload.go is doing
func main() {
	os.Clearenv() // remove all vars

	pwd, _ := os.Getwd()
	file := path.Join(pwd, ".env")
	env.ReadEnv(file)

	fmt.Println(os.Environ())
}
