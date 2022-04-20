package player

import "github.com/seniorescobar/design-patterns/cmd/flyweight/player/jersey"

type Player struct {
	Name   string
	Jersey jersey.Jersey
}
