package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

type config struct {
	numTimes   int
	printUsage bool
}

func printUsage(w io.Writer) {
	var usageString = fmt.Sprintf(`Usage: %s <integer> [-h|--help]`, os.Args[0])
	fmt.Fprintf(w, usageString)
}

func validateArgs(c config) error {
	if !(c.numTimes > 0) {
		return errors.New("Must specify a number greater than 0")
	}
	return nil
}

func parseArgs(args []string) (config, error) {
	var numTimes int
	var err error
	c := config{}
	if len(args) != 1 {
		return c, errors.New("Invalid number of arguments")
	}

	if args[0] == "-h" || args[0] == "--help" {
		c.printUsage = true
		return c, nil
	}

	numTimes, err = strconv.Atoi(args[0])
	if err != nil {
		return c, err
	}
	c.numTimes = numTimes
	return c, nil
}
