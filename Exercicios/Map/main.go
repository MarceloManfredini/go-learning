package main

import "fmt"

func main() {
	colors := map[string]string{
		"vermelho": "#ff0000",
		"verde":    "#00ff00",
		"azul":     "#0000ff",
	}
	//colors["branco"] = "#ffffff"
	//delete(colors, "vermelho")
	//fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
