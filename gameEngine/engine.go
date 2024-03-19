package engine

import (
	"tictactoe/game"
)

type Engine struct {
}

func (e Engine) GameState(board [3][3]game.Piece, players []game.Player) game.State {
	var winner game.Player

	// Check rows
	for row := 0; row < 3; row++ {
		if board[row][0] != game.Empty && board[row][0] == board[row][1] && board[row][1] == board[row][2] {
			winnerPiece := board[row][0]
			for _, player := range players {
				if player.Piece == winnerPiece {
					winner = player
					break
				}
			}
			return game.State{GameCondition: game.Winner, Winner: winner}
		}
	}

	// Check columns
	for col := 0; col < 3; col++ {
		if board[0][col] != game.Empty && board[0][col] == board[1][col] && board[1][col] == board[2][col] {
			winnerPiece := board[0][col]
			for _, player := range players {
				if player.Piece == winnerPiece {
					winner = player
					break
				}
			}
			return game.State{GameCondition: game.Winner, Winner: winner}
		}
	}

	// Check diagonals
	if board[0][0] != game.Empty && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		winnerPiece := board[0][0]
		for _, player := range players {
			if player.Piece == winnerPiece {
				winner = player
				break
			}
		}
		return game.State{GameCondition: game.Winner, Winner: winner}
	}
	if board[0][2] != game.Empty && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		winnerPiece := board[0][2]
		for _, player := range players {
			if player.Piece == winnerPiece {
				winner = player
				break
			}
		}
		return game.State{GameCondition: game.Winner, Winner: winner}
	}

	// Check for in progress
	for row := 0; row < 3; row++ {
		for col := 0; col < 3; col++ {
			if board[row][col] == game.Empty {
				return game.State{GameCondition: game.InProgress}
			}
		}
	}

	return game.State{GameCondition: game.Tie}
}
