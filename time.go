package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	utc := time.Date(2024, time.May, 17, 15, 20, 05, 0, time.UTC)
	fmt.Println(utc)
	fmt.Println(utc.Local())

	format := "2006-01-02 15:04:05"

	value := "2024-10-05 15:20:05"
	valueTime, err := time.Parse(format, value)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(valueTime)
	}

	fmt.Println(valueTime.Year())
	fmt.Println(valueTime.Month())
	fmt.Println(valueTime.Day())
	fmt.Println(valueTime.Hour())
	fmt.Println(valueTime.Minute())
	fmt.Println(valueTime.Second())
	fmt.Println(valueTime.Nanosecond())
	fmt.Println(valueTime.Location())

}
