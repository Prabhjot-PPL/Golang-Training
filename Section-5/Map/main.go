package main

import "fmt"

func main() {

	// var colors map[string]string

	// colors := make(map[int]string)
	// colors[10] = "#ffffff"

	// a map with all the keys of type string and values of string string as well
	colors := map[string]string{
		"red":   "#ff0000",
		"white": "#ffffff",
		"black": "000000",
	}

	// delete(colors, "black")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for ", color, " is ", hex)
	}
}
