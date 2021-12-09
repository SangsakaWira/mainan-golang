package main

import (
	"fmt"

	say "github.com/SangsakaWira/say-yes-module"
)

func addition(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("Hello World!")
	angka := 10
	fmt.Println(angka)
	hasil := addition(5, 6)
	fmt.Println(hasil)
	nama := say.SayYes("Sangsaka Wira")
	fmt.Println(nama)

	// ARRAY
	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	var bukus = [225]string{}

	bukus[0] = "mantap"
	bukus[1] = "mantapjiwa"

	for _, buku := range bukus {
		if buku != "" {
			fmt.Printf("nama buku : %s\n", buku)
		}
	}

	for i, fruit := range fruits {
		fmt.Printf("nama buah %d : %s\n", i, fruit)
	}

	for _, fruit := range fruits {
		fmt.Printf("nama buah : %s\n", fruit)
	}

	// SLICE
	var cars = []string{"toyota", "mazda", "mitsubishi", "tesla"}
	fmt.Println(cars[3])
}
