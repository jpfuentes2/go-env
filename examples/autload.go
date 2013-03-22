package main

import (
	"fmt"
	_ "github.com/jpfuentes2/env.go/env/autoload"
	"os"
)

func main() {
	os.Clearenv() // remove all vars
	fmt.Println(os.Environ())
}
