package main

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	expectedUsageMessage = `Usage of bin:
  -n int
    	find nearest
  -s string
    	comma separated values for the chart
`
)

func TestRun(t *testing.T) {
	t.Run("prints a help message when no args are provided", func(t *testing.T) {
		stderr := new(strings.Builder)
		err := run([]string{"bin"}, ioutil.Discard, stderr)
		assert.Equal(t, expectedUsageMessage, stderr.String())
		assert.Error(t, err, "should return an error")
	})
	t.Run("runs the correct answer without errors with minimal flags", func(t *testing.T) {
		stderr := new(strings.Builder)
		stdout := new(strings.Builder)
		err := run([]string{"bin", "-s", "1,2,3", "-n", "4"}, stdout, stderr)
		assert.Empty(t, stderr.Len(), "nothing should be written to stderr")
		assert.Equal(t, "Nearest temp: 3", stdout.String())
		assert.NoError(t, err)
	})
	t.Run("runs without errors with minimal flags and whitespace inbetween values", func(t *testing.T) {
		stderr := new(strings.Builder)
		stdout := new(strings.Builder)
		err := run([]string{"bin", "-s", "1, 2, 3", "-n", "4"}, stdout, stderr)
		assert.Empty(t, stderr.Len(), "nothing should be written to stderr")
		assert.Equal(t, "Nearest temp: 3", stdout.String())
		assert.NoError(t, err)
	})
	t.Run("returns an error when the set of chart values is empty", func(t *testing.T) {
		stderr := new(strings.Builder)
		stdout := new(strings.Builder)
		err := run([]string{"bin", "-s", "", "-n", "4"}, stdout, stderr)
		assert.Error(t, err, "should return an error")
	})
	t.Run("returns an error when a set value is not a valid integer", func(t *testing.T) {
		stderr := new(strings.Builder)
		stdout := new(strings.Builder)
		err := run([]string{"bin", "-s", "1,2,*", "-n", "4"}, stdout, stderr)
		assert.Error(t, err, "should return an error")
	})
}
