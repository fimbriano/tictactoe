package main

import (
    "tictactoe/gameEngine"
    "tictactoe/game"
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

    gs := g.CheckGameState()
    println(gs.String())
}
