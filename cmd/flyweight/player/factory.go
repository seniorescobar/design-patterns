package player

import "github.com/seniorescobar/design-patterns/cmd/flyweight/player/jersey"

type PlayerFactory struct {
	jerseyFactory *jersey.Factory
}

func NewPlayerFactory(jf *jersey.Factory) *PlayerFactory {
	return &PlayerFactory{
		jerseyFactory: jf,
	}
}

func (f *PlayerFactory) NewPlayer(name string, jerseyType jersey.JerseyType) *Player {
	return &Player{
		Name:   name,
		Jersey: f.jerseyFactory.GetJersey(jerseyType),
	}
}
