# Minesweeper Go

A classic Minesweeper game implementation in Go, played directly in your terminal with an intuitive command-line interface.

## Overview

This is a terminal-based Minesweeper game where you try to reveal all safe cells on a grid without hitting any bombs. The game features emoji-based visual representation and supports customizable board sizes.

## Features

- Interactive command-line gameplay
- Customizable square board sizes (minimum 2x2)
- Random bomb placement (50% of cells contain bombs)
- Visual feedback with emojis:
  - ⬜ Unrevealed cell
  - ⭕ Safe revealed cell
  - ❌ Bomb location (shown after game over)
- Replay functionality
- Input validation

## Prerequisites

- Go 1.25.3 or higher

## Installation

1. Clone the repository:
```bash
git clone https://github.com/iqbalnzls/minesweeper-go.git
cd minesweeper-go
```

2. Build the game:
```bash
go build -o minesweeper
```

3. Run the game:
```bash
./minesweeper
```

Alternatively, run directly without building:
```bash
go run .
```

## How to Play

1. **Start the game**: Run the executable

2. **Set board size**: Enter the dimensions for your board
   ```
   Enter row & column: 5 5
   ```
   Note: The board must be square (rows = columns) and at least 2x2

3. **Make your moves**: Enter coordinates to reveal cells
   ```
   Enter input: 0 0
   ```
   Coordinates are zero-indexed (0 to N-1 for an NxN board)

4. **Win condition**: Reveal all safe cells without hitting a bomb
   - You need to reveal exactly N safe cells to win (where N is the board size)

5. **Game over**: If you hit a bomb, the game ends and reveals all bomb locations

6. **Play again**: After each game, choose to play again (y) or quit (n)

## Game Rules

- The board is always square (NxN)
- 50% of the cells contain bombs
- You cannot click the same cell twice
- You win by revealing N safe cells (equal to the board dimension)
- Hitting a bomb ends the game immediately
- Coordinates must be within the board boundaries (0 to N-1)

## Project Structure

```
minesweeper-go/
├── main.go           # Entry point with game loop
├── minesweeper.go    # Core game logic and board functions
├── go.mod            # Go module definition
└── README.md         # This file
```

## Code Overview

### main.go
Contains the main game loop that allows players to play multiple rounds.

### minesweeper.go
Core game functionality including:
- `Position`: Struct for cell coordinates (minesweeper.go:8)
- `playMinesweeper()`: Main game flow controller (minesweeper.go:12)
- `setupBombs()`: Random bomb placement logic (minesweeper.go:67)
- `printInitialBoard()`: Display empty board (minesweeper.go:84)
- `printGameBoard()`: Display current game state (minesweeper.go:93)
- `printFinalBoard()`: Display final board with bomb locations (minesweeper.go:108)

## Example Gameplay

```
Enter row & column: 3 3
⬜ ⬜ ⬜
⬜ ⬜ ⬜
⬜ ⬜ ⬜

Enter input: 0 0
⭕ ⬜ ⬜
⬜ ⬜ ⬜
⬜ ⬜ ⬜

Enter input: 1 1
⭕ ⬜ ⬜
⬜ ⭕ ⬜
⬜ ⬜ ⬜

Enter input: 2 2
⭕ ⬜ ⬜
⬜ ⭕ ⬜
⬜ ⬜ ⭕

Congratulations!!
Play again?y/n
```

## Contributing

Contributions are welcome! Here are some ideas for improvements:
- Add bomb count indicators for adjacent cells
- Implement difficulty levels
- Add flag/mark functionality
- Support non-square boards
- Add a scoring system
- Implement save/load game state

## License

This project is open source and available under the MIT License.

## Author

[iqbalnzls](https://github.com/iqbalnzls)

## Acknowledgments

Inspired by the classic Minesweeper game originally designed by Robert Donner and Curt Johnson for Microsoft.
