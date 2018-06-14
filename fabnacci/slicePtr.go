package main

import "fmt"

func main() {
	list := make([]int, 10)
	list[0] = 0
	list[1] = 1
	fmt.Println(list)
	listPtr := &list
	fmt.Println(listPtr)

	//type *[]int does not support indexing
	//listPtr[2]=2
	(*listPtr)[2] = 2

	fmt.Println(*listPtr)
}
