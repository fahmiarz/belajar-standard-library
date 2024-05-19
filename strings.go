package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println(strings.Contains("Fahmi Arzalega", "Fahmi"))
	fmt.Println(strings.Split("Fahmi Arzalega", " "))
	fmt.Println(strings.ToLower("Fahmi Arzalega"))
	fmt.Println(strings.ToUpper("Fahmi Arzalega"))
	fmt.Println(strings.Trim("............Fahmi Arzalega.........", "."))
	fmt.Println(strings.ReplaceAll("Fahmi Arzalega Fahmi Seazz", "Fahmi", "Tehyung"))
}
