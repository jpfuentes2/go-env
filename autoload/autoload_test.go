package autoload

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func init() {
	if err := os.Chdir("../test"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func clear() {
	os.Clearenv()
}

func Test_ReadEnv(t *testing.T) {
	clear()
	Load()

	// settings in test/.env
	assert.Equal(t, "jpfuentes2", os.Getenv("TWITTER"))

	// did test/.env.local override?
	assert.Equal(t, "good", os.Getenv("LOCAL_OVERRIDE"))
}
