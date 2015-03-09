package env

import (
	"github.com/stretchr/testify/assert"
	"os"
	"syscall"
	"testing"
)

func clear() {
	os.Clearenv()
}

func Test_ReadEnv(t *testing.T) {
	clear()

	err := ReadEnv("test/.env")
	assert.NoError(t, err)
	assert.Equal(t, "jpfuentes2", os.Getenv("TWITTER"))
	assert.Equal(t, "jpfuentes2", os.Getenv("GITHUB"))
	assert.Equal(t, "http://jacquesfuentes.com", os.Getenv("WEB"))

	// should not exist
	_, found := syscall.Getenv("A_COMMENT")
	assert.False(t, found)

	// should exist but be blank
	blank, found := syscall.Getenv("BLANK_COMMENT")
	assert.True(t, found)
	assert.Equal(t, "", blank)
}

func Test_ReadEnv_DoesNotExist(t *testing.T) {
	clear()
	assert.Error(t, ReadEnv("test/adlkfjalsdjfk"))
}
