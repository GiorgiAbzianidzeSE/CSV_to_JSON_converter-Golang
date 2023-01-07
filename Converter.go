package main

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

func main() {
	fileCSV, err := os.Open("table.csv")

	if err != nil {
		panic(err)
	}

	defer fileCSV.Close()

	fileJSON, err := os.Create("table.json")

	if err != nil {
		panic(err)
	}

	defer fileJSON.Close()

	rows, err := csv.NewReader(fileCSV).ReadAll()

	if err != nil {
		panic(err)
	}
	records := make([]Record, len(rows))

	for i, v := range rows {
		splitted := strings.Split(v[0], " ")
		ageInCSV, _ := strconv.Atoi(splitted[2])
		records[i] = Record{
			FirstName: splitted[0],
			LastName:  splitted[1],
			Age:       ageInCSV,
		}
	}

	err = json.NewEncoder(fileJSON).Encode(records)

	if err != nil {
		panic(err)
	}

}

type Record struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}
