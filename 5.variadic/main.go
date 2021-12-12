package main

import "fmt"

func addition(numbers ...int) int {
	count := 0
	for _, number := range numbers {
		count += number
	}
	return count
}

func main() {
	hasil := addition(5, 5, 5, 5)
	fmt.Println(hasil)
}
