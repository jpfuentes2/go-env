package autoload

/*
  To use autoload, import the package like so:
      import _ "github.com/jpfuentes2/go-env/autoload"

  Importing this package will automatically source the `pwd`/.env file and set
  each environment variable.

  If you do not want to automatically set the env variables and/or specify a different path,
  then you should import the base "env" package and call it with a specific path. See env.go
  for more information.
*/

import (
	"github.com/jpfuentes2/go-env"
	"os"
	"path"
)

// Auto-loads `pwd`/.env
func init() {
	pwd, _ := os.Getwd()
	file := path.Join(pwd, ".env")
	env.ReadEnv(file)
}
