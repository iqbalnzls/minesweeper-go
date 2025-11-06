package main

import (
	"fmt"
	"math/rand"
)

type Position struct {
	row, col int
}

func playMinesweeper() {
	row, col := 0, 0

	fmt.Print("Enter row & column: ")
	fmt.Scanf("%d %d", &row, &col)

	if row != col || row <= 1 || col <= 1 {
		fmt.Println("Please input valid board size!!")
		return
	}

	bombs := make(map[Position]struct{})
	setupBombs(row, col, bombs)

	printInitialBoard(row, col)

	rowPos, colPos := 0, 0
	clicked := make(map[Position]struct{})
	for attempt := 0; attempt < row; {
		fmt.Print("Enter input: ")
		fmt.Scanf("%d %d", &rowPos, &colPos)

		if rowPos >= row || colPos >= col {
			fmt.Println("Please input valid value!!")
			continue
		}

		pos := Position{rowPos, colPos}

		// check if the row and column position has been inputted previously
		_, ok := clicked[pos]
		if ok {
			fmt.Println("Value has not been valid!!")
			continue
		}

		// check bombCount
		_, ok = bombs[pos]
		if ok {
			fmt.Println("BOM!!")
			printFinalBoard(row, col, clicked, bombs)

			return
		}

		clicked[pos] = struct{}{}
		printGameBoard(row, col, clicked)

		attempt++

	}

	fmt.Println("Congratulations!!")
}

func setupBombs(row, col int, bombs map[Position]struct{}) {
	bombCount := (row * col) / 2

	for {
		r, c := rand.Intn(row), rand.Intn(col)
		pos := Position{r, c}
		_, ok := bombs[pos]
		if !ok {
			bombs[pos] = struct{}{}
		}

		if len(bombs) == bombCount {
			break
		}
	}
}

func printInitialBoard(row, col int) {
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Print("⬜ ")
		}
		fmt.Println("")
	}
}

func printGameBoard(row, col int, clicked map[Position]struct{}) {
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			pos := Position{i, j}
			_, ok := clicked[pos]
			if ok {
				fmt.Print("⭕ ")
				continue
			}
			fmt.Print("⬜ ")
		}
		fmt.Println("")
	}
}

func printFinalBoard(row, col int, clicked, bombs map[Position]struct{}) {
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			pos := Position{i, j}
			_, ok := clicked[pos]
			if ok {
				fmt.Print("⭕ ")
				continue
			}

			_, ok = bombs[pos]
			if ok {
				fmt.Print("❌ ")
				continue
			}
			fmt.Print("⬜ ")
		}
		fmt.Println("")
	}
}
