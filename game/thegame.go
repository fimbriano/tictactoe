package game

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Piece string

const (
    X     Piece = "x"
    O     Piece = "o"
    Empty       = ""
)

type Condition int

const (
    InProgress Condition = iota
    Winner
    Tie
)

type Player struct {
    Name  string // name of the player
    Piece Piece  // player piece x or o
}

type State struct {
    GameCondition Condition
    Winner        Player
}

func (g State) String() string {
    switch g.GameCondition {
    case InProgress:
        return "game is still in progress"
    case Winner:
        return g.Winner.Name + " has won"
    case Tie:
        return "the game is a tie"
    default:
        return "unknown game state"
    }
}

type Engine interface {
    GameState([3][3]Piece, []Player) State
}

type TheGame struct {
    Board     [3][3]Piece
    Players   []Player
    PlayerTurn Player
    TheEngine Engine
}

func (g *TheGame) DrawBoard() {
    for i, row := range g.Board {
		for j, square := range row {
			fmt.Print(" ")
			switch square {
			case Empty:
				fmt.Print(" ")
			case X:
				fmt.Print("X")
			case O:
				fmt.Print("O")
			}
			if j != len(row)-1 {
				fmt.Print(" |")
			}
		}
		if i != len(g.Board)-1 {
			fmt.Print("\n------------")
		}
		fmt.Print("\n")
	}
}

func (g *TheGame) PlaceMark(col int, row int) {
    var piece = g.PlayerTurn.Piece;

    // Check if mark already there
    if (g.Board[row][col] != Empty) {
        fmt.Println("Piece already there. Select again.")
        g.PlayerMove()
        return
    }

    // Place mark
    g.Board[row][col] = piece
}

func (g *TheGame) PlayerMove() {
    var x, y int
    validInput := false
    scanner := bufio.NewScanner(os.Stdin)

    for !validInput {
        fmt.Println("Select a move (enter the column and row number i.e., 1 0):")
        if scanner.Scan() {
            input := scanner.Text()
            coordinates := strings.Fields(input)
            if len(coordinates) == 2 {
                if col, err := strconv.Atoi(coordinates[0]); err == nil {
                    if row, err := strconv.Atoi(coordinates[1]); err == nil {
                        if col >= 0 && col <= 2 && row >= 0 && row <= 2 {
                            x, y = col, row
                            validInput = true
                        }
                    }
                }
            }
        }

        if !validInput {
            fmt.Println("Invalid input. Please enter two numbers separated by a space (col# row#).")
        }
    }

    // Place marker
    g.PlaceMark(x, y)
}

func (g *TheGame) NextTurn() {
    if g.PlayerTurn == g.Players[0] {
        g.PlayerTurn = g.Players[1]
    } else {
        g.PlayerTurn = g.Players[0]
    }
}
func (g *TheGame) CheckGameState() State {
    return g.TheEngine.GameState(g.Board, g.Players)
}
