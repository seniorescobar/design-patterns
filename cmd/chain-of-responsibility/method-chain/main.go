package main

import (
	"fmt"

	"github.com/seniorescobar/design-patterns/cmd/chain-of-responsibility/method-chain/creature"
)

func main() {
	goblin := creature.NewCreature("Goblin", 10, 10)
	fmt.Println(goblin)

	mod := creature.NewDoubleAttackModifier(goblin)
	mod.Add(creature.NewTripleDefenseModifier(goblin))
	mod.Apply()

	fmt.Println(goblin)
}
