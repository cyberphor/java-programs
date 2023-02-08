package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
)

func getName(r io.Reader, w io.Writer) (string, error) {
	msg := "What is your name?\n"
	fmt.Fprintf(w, msg)
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	name := scanner.Text()
	if len(name) == 0 {
		return "", errors.New("Invalid name")
	}
	return name, nil
}

func greetUser(c config, name string, w io.Writer) {
	msg := fmt.Sprintf("Hello %s", name)
	for i := 0; i < c.numTimes; i++ {
		fmt.Fprintf(w, msg)
	}
}

func runCmd(r io.Reader, w io.Writer, c config) error {
	if c.printUsage {
		printUsage(w)
		return nil
	}

	name, err := getName(r, w)
	if err != nil {
		return err
	}

	greetUser(c, name, w)
	return nil
}
