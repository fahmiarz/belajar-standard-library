package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile(`e([a-z])o`)
	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("edo"))
	fmt.Println(regex.MatchString("eKo"))
	fmt.Println(regex.FindAllString("eko ego egi edo e1o eto eKo", 10))
}
