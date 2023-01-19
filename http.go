package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Table struct {
	Width  int
	Height int
}

var table = [12][12]byte{}

func TableJSON() string {
	var convertedJSON, _ = json.Marshal(table)
	return string(convertedJSON)
}

func CreateTable(t Table) {
	finalTable := [12][12]byte{}

	finalTable[0][0] = 1
	finalTable[3][0] = 1

	table = finalTable

	fmt.Println(TableJSON())
}

func TableResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, TableJSON())
}

func main() {
	fmt.Println("Testing")
	CreateTable(Table{Width: 24, Height: 24})

	http.HandleFunc("/", TableResponse)

	//Use the default DefaultServeMux.
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
