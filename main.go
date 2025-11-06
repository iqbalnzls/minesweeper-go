package main

import "fmt"

func main() {
	for {
		row, col := 0, 0

		fmt.Print("Enter row & column: ")
		fmt.Scanf("%d %d", &row, &col)

		if row != col || row <= 1 {
			fmt.Println("Please input valid board size!!")
			return
		}

		level := 0
		fmt.Println("Level:\n1:Easy; 2:Medium; 3:Hard")
		fmt.Print("Enter level: ")
		fmt.Scanf("%d", &level)

		if level < 1 || level > 3 {
			fmt.Println("Please input valid level!!")
			continue
		}

		Minesweeper(row, col, level)

		again := ""
		fmt.Println("Play again?y/n")
		fmt.Scanf("%s", &again)

		if again == "n" {
			break
		}
	}
}
