package main

import "fmt"

// zad 17
func main() {
	zad := []int{1, 2, 3, 4, 5}

	fmt.Println(Convert(zad))
}

func Convert(slice []int) []interface{} {
	inter := make([]interface{}, len(slice))
	for i, v := range slice {
		inter[i] = v
	}
	return inter
}
