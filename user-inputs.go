package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func getName() error {
	fmt.Println("What is your name?")

	reader := bufio.NewReader((os.Stdin))
	nameFromReader, err := reader.ReadString('\n')

	if err != nil {
		return err
	}

	var trimmedName string = strings.TrimSuffix(nameFromReader, "\n")

	if len(trimmedName) < 1 {
		return errors.New("Please enter a name with at least one character.")
	}

	fmt.Println("Thanks for answering, " + trimmedName + "!")

	return nil
}

func main() {
	// handle errors and exit via 1 upon error
	if err := getName(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
