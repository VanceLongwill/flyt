package flyt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	chart Chart
	temp  int
	out   int
}

func TestChart(t *testing.T) {
	testCases := []testCase{
		{
			chart: NewChart(99, -2, 0),
			temp:  9,
			out:   0,
		},
		{
			chart: NewChart(10, -10, -20),
			temp:  -18,
			out:   -20,
		},
		{
			chart: NewChart(10),
			temp:  -18,
			out:   10,
		},
	}

	for _, tx := range testCases {
		t.Run(fmt.Sprintf("should find the nearest given chart %v and input %d", tx.chart, tx.temp), func(t *testing.T) {
			assert.Equal(t, tx.out, tx.chart.NearestTemp(tx.temp))
		})

	}
}
