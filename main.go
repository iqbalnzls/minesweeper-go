package main

import (
	"fmt"
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
