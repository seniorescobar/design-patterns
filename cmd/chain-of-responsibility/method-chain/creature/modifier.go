package creature

type CreatureModifier interface {
	Add(CreatureModifier)
	Apply()
}
