package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	index := 2

	for i := 0; i < len(slice); i++ {
		slice[i-1] = slice[i]
	}
	slice = slice[:len(slice)-1]

	fmt.Println(slice)
}
