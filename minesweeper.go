package main

import (
	"fmt"
	"math/rand"
)

type Position struct {
	row, col int
}

type GameLevel float64

const (
	Easy   GameLevel = 0.3
	Medium GameLevel = 0.5
	hard   GameLevel = 0.8
)

func Minesweeper(row, col, level int) {
	bombCount, maxAttempt := calculateBombCountAndMaxAttempt(row, level)

	bombs := make(map[Position]struct{})
	setupBombs(row, col, bombCount, bombs)

	printInitialBoard(row, col)

	rowPos, colPos := 0, 0
	clicked := make(map[Position]struct{})

	for attempt := 0; attempt < maxAttempt; {
		fmt.Print("Enter input: ")
		fmt.Scanf("%d %d", &rowPos, &colPos)
		rowPos, colPos = rowPos-1, colPos-1

		if rowPos >= row || colPos >= col {
			fmt.Println("Please input valid value!!")
			continue
		}

		pos := Position{rowPos, colPos}

		// check if the row and column position has been inputted previously
		_, ok := clicked[pos]
		if ok {
			fmt.Println("Position has been inputted previously!!")
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

func calculateBombCountAndMaxAttempt(row, level int) (bombCount, maxAttempt int) {
	switch level {
	case 1:
		bombCount = int(float64(row*row) * float64(Easy))
		maxAttempt = int(float64(row*row-bombCount) * float64(Easy))
	case 2:
		bombCount = int(float64(row*row) * float64(Medium))
		maxAttempt = int(float64(row*row-bombCount) * float64(Medium))
	case 3:
		bombCount = int(float64(row*row) * float64(hard))
		maxAttempt = int(float64(row*row-bombCount) * float64(hard))
	}

	if maxAttempt < 1 {
		maxAttempt = 1
	}
	return
}

func setupBombs(row, col, bombCount int, bombs map[Position]struct{}) {
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
