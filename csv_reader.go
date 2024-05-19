package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "Fahmi,Arzalega,Seazz\n" +
		"Budi,Anugraha,Jaya\n" +
		"Odetta,Firstly,Nourma\n"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}
}
