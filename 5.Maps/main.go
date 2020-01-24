package main

import "fmt"

func main() {
	// Maps are dictionaries, with a typed key and value
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"white": "#ffffff",
	}

	// Can be added to
	colors["black"] = "#000000"
	printMap(colors)

	// And removed
	delete(colors, "black")
	printMap(colors)
}

func printMap(colors map[string]string) {
	for k, v := range colors {
		fmt.Println("Hex code for", k, "is", v)
	}
	fmt.Println()
}
