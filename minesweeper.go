package main

import (
	"fmt"
	"math/rand"
	"slices"
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

func singlePlayer() {
	trackList := make([]*PlayerTrack, 0)

	name := ""
	fmt.Print("Player name: ")
	fmt.Scanf("%s", &name)

	for {
		row, col := 0, 0

		fmt.Print("Enter row & column: ")
		fmt.Scanf("%d %d", &row, &col)

		if row != col || row <= 1 {
			fmt.Println("Please input valid board size!!")
			continue
		}

		level := 0
		fmt.Println("Level:\n1:Easy; 2:Medium; 3:Hard")
		fmt.Print("Enter level: ")
		fmt.Scanf("%d", &level)

		if level < 1 || level > 3 {
			fmt.Println("Please input valid level!!")
			continue
		}

		track := playMinesweeper(row, col, level)
		track.Name = name
		trackList = append(trackList, track)

		again := ""
		fmt.Println("Play again?y/n")
		fmt.Scanf("%s", &again)

		if again == "n" {
			break
		}
	}

	fmt.Printf("===============\n%s Game Records:\n", name)
	if len(trackList) == 0 {
		fmt.Println("All round were failed")
		return
	}

	for i, track := range trackList {
		fmt.Printf("Round: %2d | Success: %-5v | Board size: %2dx%-2d | Level: %-6s | Duration: %v\n", i+1, !track.IsFailed, track.BoardSize, track.BoardSize, GameLevel[track.Level], track.Duration)
	}
}

func multiplayer() {
	trackList := make([]*PlayerTrack, 0)
	for {
		row, col := 0, 0

		name := ""
		fmt.Print("Player name: ")
		fmt.Scanf("%s", &name)

		fmt.Print("Enter row & column: ")
		fmt.Scanf("%d %d", &row, &col)

		if row != col || row <= 1 {
			fmt.Println("Please input valid board size!!")
			continue
		}

		level := 0
		fmt.Println("Level:\n1:Easy; 2:Medium; 3:Hard")
		fmt.Print("Enter level: ")
		fmt.Scanf("%d", &level)

		if level < 1 || level > 3 {
			fmt.Println("Please input valid level!!")
			continue
		}

		track := playMinesweeper(row, col, level)
		track.Name = name
		trackList = append(trackList, track)

		again := ""
		fmt.Println("Play again?y/n")
		fmt.Scanf("%s", &again)

		if again == "n" {
			break
		}
	}

	sortLeaderboard(trackList)

	fmt.Println("\n===============\nLeaderboard:")
	if len(trackList) == 0 {
		fmt.Println("All player were failed")
		return
	}

	for _, track := range trackList {
		fmt.Printf("Name: %-12s | Success: %-5v | Board size: %2dx%-2d | Level: %-6s | Duration: %v\n", track.Name, !track.IsFailed, track.BoardSize, track.BoardSize, GameLevel[track.Level], track.Duration)
	}
}

func sortLeaderboard(trackList []*PlayerTrack) {
	slices.SortFunc(trackList, func(a, b *PlayerTrack) int {
		// First priority: Passed = true comes before Passed = false
		if !a.IsFailed && b.IsFailed {
			return -1 // a comes first
		}
		if a.IsFailed && !b.IsFailed {
			return 1 // b comes first
		}

		// Second priority: check player level
		if a.Level > b.Level {
			return -1 // a come first (harder level)
		}
		if a.Level < b.Level {
			return 1 // b come first
		}

		// Third priority: check by board size
		if a.BoardSize > b.BoardSize {
			return -1
		}
		if a.BoardSize < b.BoardSize {
			return 1
		}

		// Fourth priority: If both have the same Passed status, sort by Duration
		if a.Duration < b.Duration {
			return -1 // a comes first (less duration)
		}
		if a.Duration > b.Duration {
			return 1 // b comes first
		}

		return 0
	})
}

func playMinesweeper(row, col, level int) *PlayerTrack {
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
