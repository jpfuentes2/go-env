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

// automatically load the env
func init() {
	Load()
}

// Load loads `pwd`/.env or GOENV and a `.local` for overrides
func Load() {
	pwd, _ := os.Getwd()
	goenv := os.Getenv("GOENV")
	envFile := ".env"

	if len(goenv) > 0 && goenv != "development" {
		envFile = fmt.Sprintf("%s.%s", envFile, goenv)
	}

	localFile := envFile + ".local"
	env.ReadEnv(path.Join(pwd, envFile))
	env.ReadEnv(path.Join(pwd, localFile))
}
