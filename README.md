# Minesweeper Go

A classic Minesweeper game implementation in Go, played directly in your terminal with an intuitive command-line interface.

## Overview

This is a terminal-based Minesweeper game where you try to reveal all safe cells on a grid without hitting any bombs. The game features emoji-based visual representation, customizable board sizes, and competitive time-based gameplay with a 10-second per-move limit.

Choose between two game modes:
- **Single Player**: Practice mode where you can play multiple rounds and track your performance across rounds
- **Multiplayer**: Competitive mode with a comprehensive leaderboard system that ranks players by board size, difficulty level, and completion time

## Features

- Interactive command-line gameplay
- Two game modes:
  - **Single Player**: Practice mode with round tracking
  - **Multiplayer**: Competitive mode with leaderboard rankings
- Customizable square board sizes (minimum 2x2)
- Three difficulty levels with unified percentage system (30%, 50%, 80%)
- Random bomb placement
- Visual feedback with emojis:
  - ⬜ Unrevealed cell
  - ⭕ Safe revealed cell
  - ❌ Bomb location (shown after game over)
- Real-time game timer tracking
- 10-second time limit per move
- Leaderboard/ranking system (multiplayer mode)
- Round-by-round records (single player mode)
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

2. **Select game mode**: Choose between single or multiplayer
   ```
   Select gameplay
   1:Single player; 2:Multi player
   ```

### Single Player Mode

3. **Enter player name**: Provide your name once for all rounds
   ```
   Player name: alice
   ```

4. **Set board size**: Enter the dimensions for your board
   ```
   Enter row & column: 5 5
   ```
   Note: The board must be square (rows = columns) and at least 2x2

5. **Select difficulty level**: Choose your preferred difficulty
   ```
   Level:
   1:Easy; 2:Medium; 3:Hard
   Enter level: 2
   ```
   - **Easy (1)**: 30% difficulty (30% bombs, reveal 30% of safe cells to win)
   - **Medium (2)**: 50% difficulty (50% bombs, reveal 50% of safe cells to win)
   - **Hard (3)**: 80% difficulty (80% bombs, reveal 80% of safe cells to win)

6. **Make your moves**: Enter coordinates to reveal cells
   ```
   Enter input: 1 1
   ```
   - Coordinates start from 1 (1 to N for an NxN board)
   - **Important**: You have a 10-second time limit per move. Exceeding this results in game failure.

7. **Win condition**: Reveal a percentage of safe cells without hitting a bomb
   - **Easy**: Reveal 30% of safe cells to win
   - **Medium**: Reveal 50% of safe cells to win
   - **Hard**: Reveal 80% of safe cells to win
   - Example: On a 10x10 board with Easy (70 safe cells), you need to reveal only 21 cells to win

8. **Game over conditions**:
   - Hit a bomb → reveals all bomb locations and you fail
   - Exceed 10 seconds on any single move → game fails
   - Invalid input is not counted against you

9. **Play again**: After each game, choose to play again (y) or quit (n)

10. **View records**: After quitting, see your round-by-round performance including:
    - Round number
    - Success/failure status
    - Board size and difficulty level
    - Completion time

### Multiplayer Mode

Same steps as single player, but with these differences:

- **Player name entry**: Enter a new player name for each game session
- **Leaderboard display**: After quitting, see:
    - Success status (successful players rank above failed players)
    - Difficulty level (harder levels rank higher)
    - Board size (larger boards rank higher)
    - Completion time (faster times rank higher)
- **All games shown**: Both successful and failed games appear with Success status indicated

## Game Rules

- The board is always square (NxN)
- Difficulty uses a **unified percentage system** - the same percentage applies to both bomb density and win condition:
  - **Easy (30%)**: 30% of cells are bombs, reveal 30% of safe cells to win
  - **Medium (50%)**: 50% of cells are bombs, reveal 50% of safe cells to win
  - **Hard (80%)**: 80% of cells are bombs, reveal 80% of safe cells to win
- You cannot click the same cell twice
- **Time limit**: Each move must be completed within 10 seconds
- Hitting a bomb ends the game immediately
- Exceeding the time limit on any move ends the game immediately
- Coordinates must be within the board boundaries (1 to N)
- Valid difficulty levels are 1 (Easy), 2 (Medium), or 3 (Hard)
- **Single player mode**: All rounds (successful and failed) are shown in the game records with round numbers
- **Multiplayer mode**: All games (successful and failed) are shown in the leaderboard with Success status
- Both modes display successful and failed games, with sorting prioritizing successful games first

### Difficulty Scaling Examples

**5x5 Board (25 cells):**
- **Easy**: 7 bombs, 18 safe cells → need 5 correct moves (30% of 18)
- **Medium**: 12 bombs, 13 safe cells → need 6 correct moves (50% of 13)
- **Hard**: 20 bombs, 5 safe cells → need 4 correct moves (80% of 5)

**10x10 Board (100 cells):**
- **Easy**: 30 bombs, 70 safe cells → need 21 correct moves (30% of 70)
- **Medium**: 50 bombs, 50 safe cells → need 25 correct moves (50% of 50)
- **Hard**: 80 bombs, 20 safe cells → need 16 correct moves (80% of 20)

## Leaderboard System (Multiplayer Mode)

The multiplayer mode features a comprehensive sorting system that displays all player attempts:

### Sorting Criteria (in priority order):
1. **Success Status**: Successful games (no bomb hits, no timeouts) rank above failed games
2. **Difficulty Level**: Harder levels rank higher (Hard > Medium > Easy)
3. **Board Size**: Larger boards rank higher (10x10 > 5x5 > 3x3)
4. **Completion Time**: Faster times rank higher within the same level and board size

### Examples:
- A successful player on 3x3 Easy ranks higher than a failed player on 10x10 Hard
- Among successful players: 5x5 Hard ranks higher than 10x10 Easy (difficulty first)
- Among successful players at same level: 10x10 Medium ranks higher than 5x5 Medium
- Two successful players completing 5x5 Hard are ranked by fastest time

### Output Sections:
- **Leaderboard**: Shows all games sorted by the criteria above

### Notes:
- **Multiplayer mode**: Shows Leaderboard sections
- **Single player mode**: Shows only Game Records with round numbers
- Both modes show all attempts (successful and failed)
- Successful games are sorted to the top of the leaderboard
- Each game session is tracked independently
- Displays at the end when you choose to quit (n)

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
Contains the game mode selection and gameplay functions:
- `main()`: Game mode selection menu (main.go:7)
- `singlePlayer()`: Practice mode with round tracking (main.go:19)
  - Player enters name once
  - Tracks all rounds in `PlayerTrack` collection
  - Displays round-by-round performance summary at the end
- `multiplayer()`: Competitive mode with leaderboard (main.go:72)
  - Each player enters their name per game
  - Tracks all players in `PlayerTrack` collection
  - Displays game records and leaderboard at the end
- `sortLeaderboard()`: Sorts leaderboard by multiple criteria (main.go:125)
  - Sorting order: success status → difficulty level → board size → completion time

### minesweeper.go
Core game functionality including:
- `Position`: Struct for cell coordinates (minesweeper.go:9)
- `PlayerTrack`: Struct for tracking player performance with fields:
  - `Name`: Player name (minesweeper.go:14)
  - `Duration`: Game completion time (minesweeper.go:15)
  - `Level`: Difficulty level (1-3) (minesweeper.go:16)
  - `IsFailed`: Whether player failed (bomb/timeout) (minesweeper.go:17)
  - `BoardSize`: Size of the board (minesweeper.go:18)
- `GameLevelPercentage`: Custom type for difficulty percentages (minesweeper.go:21)
- Difficulty constants using unified percentage system:
  - `Easy` (30%) - used for both bomb density and win threshold (minesweeper.go:24)
  - `Medium` (50%) - used for both bomb density and win threshold (minesweeper.go:25)
  - `hard` (80%) - used for both bomb density and win threshold (minesweeper.go:26)
- `GameLevel` map: Maps level numbers to difficulty names (minesweeper.go:30)
- `Minesweeper()`: Main game flow controller with time tracking and 10s limit (minesweeper.go:37)
- `calculateBombCountAndMaxAttempt()`: Calculates bombs and required moves using the same percentage (minesweeper.go:95)
- `setupBombs()`: Random bomb placement logic (minesweeper.go:114)
- `printInitialBoard()`: Display empty board (minesweeper.go:129)
- `printGameBoard()`: Display current game state (minesweeper.go:138)
- `printFinalBoard()`: Display final board with bomb locations (minesweeper.go:153)

## Example Gameplay

### Single Player Mode
```
Select gameplay
1:Single player; 2:Multi player
1

Player name: alice
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
y

Enter row & column: 4 4
[...second round gameplay...]
Congratulations!!
Play again?y/n
n

===============
alice Game Records:
Round:  1 | Success: true  | Board size:  3x3  | Level: Easy   | Duration: 8s
Round:  2 | Success: true  | Board size:  4x4  | Level: Medium | Duration: 12s
```

**Note**: Single player mode shows round-by-round records, including both successful and failed attempts.

### Multiplayer Mode with Leaderboard
```
Select gameplay
1:Single player; 2:Multi player
2

Player name: alice
Enter row & column: 5 5
Level:
1:Easy; 2:Medium; 3:Hard
Enter level: 2
[...gameplay...]
Congratulations!!
Play again?y/n
y

Player name: bob
Enter row & column: 5 5
Level:
1:Easy; 2:Medium; 3:Hard
Enter level: 3
[...gameplay...]
Congratulations!!
Play again?y/n
y

Player name: charlie
Enter row & column: 4 4
Level:
1:Easy; 2:Medium; 3:Hard
Enter level: 1
[...gameplay...]
BOM!!
Play again?y/n
n

===============
Leaderboard:
Name: bob          | Success: true  | Board size:  5x5  | Level: Hard   | Duration: 15s
Name: alice        | Success: true  | Board size:  5x5  | Level: Medium | Duration: 12s
Name: charlie      | Success: false | Board size:  4x4  | Level: Easy   | Duration: 5s
```

**Understanding the Output**:
- **Leaderboard**: Shows all games sorted by the ranking criteria:
  - Successful players (bob, alice) rank above failed players (charlie)
  - Among successful players: Bob ranks #1 because Hard difficulty ranks above Medium
  - Among players at same difficulty and board size, faster times rank higher
  - Failed players appear at the bottom with Success: false status
- All columns are properly aligned for easy reading

## Improvements

- [ ] Add bomb count indicators for adjacent cells (classic minesweeper numbers)
- [X] Implement difficulty levels with unified percentage system (Easy: 30%, Medium: 50%, Hard: 80%)
- [ ] Add flag/mark functionality to mark suspected bombs
- [ ] Support non-square boards (rectangular grids)
- [X] Implement game mode selection (Single Player / Multiplayer)
- [X] Add single player mode with round tracking
- [X] Add a leaderboard/ranking system (multiplayer mode)
- [X] Add time tracking for each game
- [X] Implement per-move time limit (10 seconds)
- [X] Add multiplayer support with player names
- [ ] Implement save/load game state
- [ ] Implement cascading reveal for cells with no adjacent bombs
- [ ] Display progress indicator (e.g., "5/18 cells revealed")
- [ ] Add best time records per difficulty/board size
- [ ] Export leaderboard to file (CSV/JSON)


## Author

[iqbalnzls](https://github.com/iqbalnzls)

## Acknowledgments

Inspired by the classic Minesweeper game originally designed by Robert Donner and Curt Johnson for Microsoft.
