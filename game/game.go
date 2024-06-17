package game

type Game struct {
	Player   *Player
	Position *Position
}

type Position struct {
	X int
	Y int
}

type Direction int

func NewGame(player *Player) *Game {
	return &Game{
		Player:   player,
		Position: &Position{},
	}
}
