package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("Fahmi")
	data.PushBack("Arzalega")
	data.PushBack("Seazzz")
	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
