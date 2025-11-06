# Minesweeper Go

A classic Minesweeper game implementation in Go, played directly in your terminal with an intuitive command-line interface.

## Overview

This is a terminal-based Minesweeper game where you try to reveal all safe cells on a grid without hitting any bombs. The game features emoji-based visual representation and supports customizable board sizes.

## Features

- Interactive command-line gameplay
- Customizable square board sizes (minimum 2x2)
- Three difficulty levels with unified percentage system (30%, 50%, 80%)
- Random bomb placement
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

3. **Select difficulty level**: Choose your preferred difficulty
   ```
   Level:
   1:Easy; 2:Medium; 3:Hard
   Enter level: 2
   ```
   - **Easy (1)**: 30% difficulty (30% bombs, reveal 30% of safe cells to win)
   - **Medium (2)**: 50% difficulty (50% bombs, reveal 50% of safe cells to win)
   - **Hard (3)**: 80% difficulty (80% bombs, reveal 80% of safe cells to win)

4. **Make your moves**: Enter coordinates to reveal cells
   ```
   Enter input: 1 1
   ```
   Coordinates start from 1 (1 to N for an NxN board)

5. **Win condition**: Reveal a percentage of safe cells without hitting a bomb
   - **Easy**: Reveal 30% of safe cells to win
   - **Medium**: Reveal 50% of safe cells to win
   - **Hard**: Reveal 80% of safe cells to win
   - Example: On a 10x10 board with Easy (70 safe cells), you need to reveal only 21 cells to win

6. **Game over**: If you hit a bomb, the game ends and reveals all bomb locations

7. **Play again**: After each game, choose to play again (y) or quit (n)

## Game Rules

- The board is always square (NxN)
- Difficulty uses a **unified percentage system** - the same percentage applies to both bomb density and win condition:
  - **Easy (30%)**: 30% of cells are bombs, reveal 30% of safe cells to win
  - **Medium (50%)**: 50% of cells are bombs, reveal 50% of safe cells to win
  - **Hard (80%)**: 80% of cells are bombs, reveal 80% of safe cells to win
- You cannot click the same cell twice
- Hitting a bomb ends the game immediately
- Coordinates must be within the board boundaries (1 to N)
- Valid difficulty levels are 1 (Easy), 2 (Medium), or 3 (Hard)

### Difficulty Scaling Examples

**5x5 Board (25 cells):**
- **Easy**: 7 bombs, 18 safe cells → need 5 correct moves (30% of 18)
- **Medium**: 12 bombs, 13 safe cells → need 6 correct moves (50% of 13)
- **Hard**: 20 bombs, 5 safe cells → need 4 correct moves (80% of 5)

**10x10 Board (100 cells):**
- **Easy**: 30 bombs, 70 safe cells → need 21 correct moves (30% of 70)
- **Medium**: 50 bombs, 50 safe cells → need 25 correct moves (50% of 50)
- **Hard**: 80 bombs, 20 safe cells → need 16 correct moves (80% of 20)

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
Contains the main game loop that allows players to play multiple rounds. Handles user input for board size, difficulty level selection, and replay functionality.

### minesweeper.go
Core game functionality including:
- `Position`: Struct for cell coordinates (minesweeper.go:8)
- `GameLevel`: Custom type for difficulty percentages (minesweeper.go:12)
- Difficulty constants using unified percentage system:
  - `Easy` (30%) - used for both bomb density and win threshold (minesweeper.go:15)
  - `Medium` (50%) - used for both bomb density and win threshold (minesweeper.go:16)
  - `hard` (80%) - used for both bomb density and win threshold (minesweeper.go:17)
- `Minesweeper()`: Main game flow controller (minesweeper.go:20)
- `calculateBombCountAndMaxAttempt()`: Calculates bombs and required moves using the same percentage (minesweeper.go:69)
- `setupBombs()`: Random bomb placement logic (minesweeper.go:88)
- `printInitialBoard()`: Display empty board (minesweeper.go:103)
- `printGameBoard()`: Display current game state (minesweeper.go:112)
- `printFinalBoard()`: Display final board with bomb locations (minesweeper.go:127)

## Example Gameplay

```
Enter row & column: 3 3
Level:
1:Easy; 2:Medium; 3:Hard
Enter level: 1
⬜ ⬜ ⬜
⬜ ⬜ ⬜
⬜ ⬜ ⬜

Enter input: 1 1
⭕ ⬜ ⬜
⬜ ⬜ ⬜
⬜ ⬜ ⬜

Enter input: 2 2
⭕ ⬜ ⬜
⬜ ⭕ ⬜
⬜ ⬜ ⬜

Congratulations!!
Play again?y/n
```

**Note**: On a 3x3 board with Easy mode (30% difficulty), there are 2 bombs and 7 safe cells. You only need to reveal 2 safe cells (30% of 7) to win!

## Contributing

Contributions are welcome! Here are some ideas for improvements:
- [ ] Add bomb count indicators for adjacent cells (classic minesweeper numbers)
- [X] Implement difficulty levels with unified percentage system (Easy: 30%, Medium: 50%, Hard: 80%)
- [ ] Add flag/mark functionality to mark suspected bombs
- [ ] Support non-square boards (rectangular grids)
- [ ] Add a scoring system (time-based or moves-based)
- [ ] Implement save/load game state
- [ ] Add a timer to track game duration
- [ ] Implement cascading reveal for cells with no adjacent bombs
- [ ] Display progress indicator (e.g., "5/18 cells revealed")

## License

This project is open source and available under the MIT License.

## Author

[iqbalnzls](https://github.com/iqbalnzls)

## Acknowledgments

Inspired by the classic Minesweeper game originally designed by Robert Donner and Curt Johnson for Microsoft.
