package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type House struct {
	Type string `validate:"required"`
	Size int    `validate:"gte=0,lte=130"`
}

var HouseList = []House{}

var pl = fmt.Println

func variables() error {
	var firstName string = "John"
	var lastName string = "Doe"
	var age int = 26

	pl("First Name:", firstName, "Last Name", lastName, ", age: ", age)

	var MyHouse House = House{Type: "Terraced", Size: 240}
	pl("My House", MyHouse)

	HouseList = append(HouseList, MyHouse)
	pl("HouseList:")
	pl(HouseList)

	fmt.Println("Enter a house type")

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
	if err := variables(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
