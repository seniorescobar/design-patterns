package creature

type DoubleAttackModifier struct {
	creature *Creature
	next     CreatureModifier
}

func NewDoubleAttackModifier(creature *Creature) CreatureModifier {
	return &DoubleAttackModifier{
		creature: creature,
	}
}

func (m *DoubleAttackModifier) Add(next CreatureModifier) {
	m.next = next
}

func (m *DoubleAttackModifier) Apply() {
	m.creature.Attack *= 2

	if m.next != nil {
		m.next.Apply()
	}
}
