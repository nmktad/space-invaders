package game

type Invader struct {
	Position *Position
	HP       int
}

type InvaderFactoryInterface interface {
	CreateInvader(Position) *Invader
}

type InvaderFactory struct{}

func (f *InvaderFactory) CreateInvader(p *Position, hp int) *Invader {
	return &Invader{
		Position: p,
		HP:       hp,
	}
}
