package main

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type userSlice []User

func (s userSlice) Len() int { //kontrak sort.interface
	return len(s)
}

func (s userSlice) Less(i, j int) bool { //kontrak sort.interface
	return s[i].Age < s[j].Age
}

func (s userSlice) Swap(i, j int) { //kontrak sort.interface
	s[i], s[j] = s[j], s[i]
}
func main() {
	users := []User{
		{"Fahmi", 24},
		{"Eko", 30},
		{"Bagas", 22},
		{"Wayan", 55},
	}

	sort.Sort(userSlice(users))
	fmt.Println(users)

}
