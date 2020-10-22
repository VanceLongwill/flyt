package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	// ErrNotEnoughArgs is returned when fewer than the required args are provided
	ErrNotEnoughArgs = errors.New("not enough args provided")
)

func run(args []string, stdout io.Writer, stderr io.Writer) error {
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	flags.SetOutput(stderr)
	var (
		verbose = flags.Bool("v", false, "verbose logging")
	)

	if err := flags.Parse(args[1:]); err != nil {
		return err
	}

	if len(args) < 2 {
		flags.Usage()
		return ErrNotEnoughArgs
	}

	if *verbose {
		fmt.Fprintln(stdout, "Verbose logging enabled")
	}

	return nil
}

func main() {
	if err := run(os.Args, os.Stdout, os.Stderr); err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	}
}
