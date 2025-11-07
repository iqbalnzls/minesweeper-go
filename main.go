package main

import (
	"fmt"
	"slices"
)

func main() {
	types := 0
	fmt.Print("Select gameplay\n1:Single player; 2:Multi player\n")
	fmt.Scanf("%d", &types)

	if types == 1 {
		singlePlayer()
	} else if types == 2 {
		multiplayer()
	}
}

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

		track := Minesweeper(row, col, level)
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

		track := Minesweeper(row, col, level)
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
