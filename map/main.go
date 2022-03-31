package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff000",
		"green": "#4bf745",
		"white": "#fffff",
	}

	//var colors map[string]string

	//this will make an empty map
	//colors := make(map[string]string)

	//colors["white"] = "#fffff"

	//delete(colors, "white")

	//fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for",color,"is",hex)
	}
}
