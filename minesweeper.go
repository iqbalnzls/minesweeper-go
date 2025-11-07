package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Position struct {
	row, col int
}

type PlayerTrack struct {
	Name      string
	Duration  time.Duration
	Level     int
	IsFailed  bool
	BoardSize int
}

type GameLevelPercentage float64

const (
	Easy   GameLevelPercentage = 0.3
	Medium GameLevelPercentage = 0.5
	hard   GameLevelPercentage = 0.8
)

var (
	GameLevel = map[int]string{
		1: "Easy",
		2: "Medium",
		3: "Hard",
	}
)

func Minesweeper(row, col, level int) *PlayerTrack {
	start := time.Now()
	bombCount, maxAttempt := calculateBombCountAndMaxAttempt(row, level)

	bombs := make(map[Position]struct{})
	setupBombs(row, col, bombCount, bombs)

	printInitialBoard(row, col)

	rowPos, colPos := 0, 0
	clicked := make(map[Position]struct{})

	for attempt := 0; attempt < maxAttempt; {
		tn := time.Now()
		fmt.Print("Enter input: ")
		fmt.Scanf("%d %d", &rowPos, &colPos)
		rowPos, colPos = rowPos-1, colPos-1

		//validate time taken by the user for each input should not the more 10s
		if time.Since(tn).Round(time.Second) > 10*time.Second {
			fmt.Println("Time limit exceeded!!")
			return &PlayerTrack{Duration: time.Since(start).Round(time.Second), Level: level, IsFailed: true, BoardSize: row}
		}

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

			return &PlayerTrack{Duration: time.Since(start).Round(time.Second), Level: level, IsFailed: true, BoardSize: row}
		}

		clicked[pos] = struct{}{}
		printGameBoard(row, col, clicked)

		attempt++

	}

	fmt.Println("Congratulations!!")
	return &PlayerTrack{Duration: time.Since(start).Round(time.Second), Level: level, IsFailed: false, BoardSize: row}
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
