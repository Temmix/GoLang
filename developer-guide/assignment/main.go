package main

import "fmt"

func main() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	printEvenAndOdd(data)
}

func printEvenAndOdd(data []int) {
	for _, value := range data {
		if value%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}
