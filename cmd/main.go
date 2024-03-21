package main

import (
	"fmt"
	"tictactoe/game"
	engine "tictactoe/gameEngine"
)

func main() {
    var g game.TheGame
    // Set the players
    g.Players = []game.Player{{
        Name:  "Felicia",
        Piece: game.X,
    }, {
        Name:  "Danielle",
        Piece: game.O,
    }}
    // Set player 1
    g.PlayerTurn = g.Players[0];
    // Set empty board
    g.Board = [3][3]game.Piece{{game.Empty, game.Empty, game.Empty}, {game.Empty, game.Empty, game.Empty}, {game.Empty, game.Empty, game.Empty}}
    // Set engine
    g.TheEngine = engine.Engine{}

    // Check if players is valid
    if len(g.Players) != 2 {
        fmt.Println("Error: There must be exactly two players.")
        return
    }
    
    for{
        fmt.Printf("%s's turn (%s)\n", g.PlayerTurn.Name, g.PlayerTurn.Piece)

        // Draw the board
        g.DrawBoard()

        // Get player move
        g.PlayerMove()

        // Check if player won
        gs := g.CheckGameState()
        if (gs.GameCondition != game.InProgress ) {
            g.DrawBoard()
            println(gs.String())
            break
        }

        // Switch player
        g.NextTurn()
    }
}
