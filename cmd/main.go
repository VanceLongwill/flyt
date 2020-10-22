package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/vancelongwill/flyt"
)

var (
	// ErrNotEnoughArgs is returned when fewer than the required args are provided
	ErrNotEnoughArgs = errors.New("not enough args provided")
)

func run(args []string, stdout io.Writer, stderr io.Writer) error {
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	flags.SetOutput(stderr)
	var (
		nearest = flags.Int("n", 0, "find nearest")
		set     = flags.String("s", "", "comma separated values for the chart")
	)

	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	if len(args) < 2 {
		flags.Usage()
		return ErrNotEnoughArgs
	}

	if nearest == nil || set == nil {
		return ErrNotEnoughArgs
	}

	setStrings := strings.Split(*set, ",")

	if len(setStrings) == 0 {
		return errors.New("set must contain at least one temperature value")
	}

	var tempSet []int

	for i, s := range setStrings {
		val, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil {
			return fmt.Errorf("unable to parse value at index = '%d': %w", i, err)
		}
		tempSet = append(tempSet, val)
	}

	chart := flyt.NewChart(tempSet...)

	fmt.Fprintf(stdout, "Nearest temp: %d", chart.NearestTemp(*nearest))

	return nil
}

func main() {
	if err := run(os.Args, os.Stdout, os.Stderr); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	}
}
