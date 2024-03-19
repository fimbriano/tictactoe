package game

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
    TheEngine Engine
}

func (g TheGame) CheckGameState() State {
    return g.TheEngine.GameState(g.Board, g.Players)
}
