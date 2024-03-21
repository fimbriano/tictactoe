package main

import (
	"fmt"
	"tictactoe/game"
	"tictactoe/gameEngine"
)

func main() {
    var g game.TheGame
    g.Players = []game.Player{{
        Name:  "Felicia",
        Piece: game.X,
    }, {
        Name:  "Danielle",
        Piece: game.O,
    }}
    g.Board = [3][3]game.Piece{{game.O, game.X, game.X}, {game.O, game.O, game.X}, {game.X, game.O, game.O}}
    g.TheEngine = engine.Engine{}

     // Check if players are valid
     if len(g.Players) != 2 {
        fmt.Println("Error: There must be exactly two players.")
        return
    }
    
    gs := g.CheckGameState()
    println(gs.String())
}
