package main

import "fmt"

func main() {
	for {
		playMinesweeper()

		again := ""

		fmt.Println("Play again?y/n")
		fmt.Scanf("%s", &again)

		if again == "n" {
			break
		}
	}
}
