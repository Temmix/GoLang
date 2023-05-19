package main

import "fmt"

func main() {
	// working with map
	colors := map[string]string{
		"red":   "#ff000",
		"green": "#fefef",
		"white": "#fffff",
	}

	fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println(key, value)
	}
}
