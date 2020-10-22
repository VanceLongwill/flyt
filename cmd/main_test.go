package main

import (
	"io/ioutil"
	"strings"
	"testing"

	// "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

var (
	expectedUsageMessage = `Usage of bin:
  -v	verbose logging
`
)

func TestRun(t *testing.T) {
	t.Run("prints a help message when no args are provided", func(t *testing.T) {
		stderr := new(strings.Builder)
		err := run([]string{"bin"}, ioutil.Discard, stderr)
		assert.Equal(t, expectedUsageMessage, stderr.String())
		assert.Error(t, err, "should return an error")
	})
	t.Run("runs without errors with minimal flags", func(t *testing.T) {
		stderr := new(strings.Builder)
		stdout := new(strings.Builder)
		err := run([]string{"bin", "-v"}, stdout, stderr)
		assert.Empty(t, stderr.Len(), "nothing should be written to stderr")
		assert.Equal(t, "Verbose logging enabled\n", stdout.String())
		assert.NoError(t, err)
	})

	// t.Run("tests a mocked interface", func(t *testing.T) {
	// 	ctrl := gomock.NewController(t)
	// 	defer ctrl.Finish()
	// })
}
