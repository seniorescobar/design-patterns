package main

import (
	"fmt"

	"github.com/seniorescobar/design-patterns/cmd/flyweight/player"
	"github.com/seniorescobar/design-patterns/cmd/flyweight/player/jersey"
)

func main() {
	pf := player.NewPlayerFactory(
		jersey.NewJerseyFactory(),
	)

	players := make([]*player.Player, 0)
	add := func(p *player.Player) {
		players = append(players, p)
	}

	// Porto
	add(pf.NewPlayer("Otavio", jersey.PortoJerseyType))
	add(pf.NewPlayer("Evanilson", jersey.PortoJerseyType))
	add(pf.NewPlayer("Taremi", jersey.PortoJerseyType))

	// Benfica
	add(pf.NewPlayer("Nunez", jersey.BenficaJerseyType))

	for _, p := range players {
		fmt.Printf("%+v\n", p)
	}
}
