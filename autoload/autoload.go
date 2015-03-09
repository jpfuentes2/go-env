/*
  To use autoload, import the package like so:
      import _ "github.com/jpfuentes2/go-env/autoload"

  Importing this package will automatically source the `pwd`/.env file and set
  each environment variable.

  If you do not want to automatically set the env variables and/or specify a different path,
  then you should import the base "env" package and call it with a specific path. See env.go
  for more information.
*/
package autoload

import (
	"fmt"
	"github.com/jpfuentes2/go-env"
	"os"
	"path"
)

// Auto-loads `pwd`/.env or GOENV
func init() {
	pwd, _ := os.Getwd()
	goenv := os.Getenv("GOENV")
	file := ".env"

	if len(goenv) > 0 && goenv != "development" {
		file = fmt.Sprintf("%s.%s", file, goenv)
	}

	env.ReadEnv(path.Join(pwd, file))
}
