package creature

type TripleDefenseModifier struct {
	creature *Creature
	next     CreatureModifier
}

func NewTripleDefenseModifier(creature *Creature) CreatureModifier {
	return &TripleDefenseModifier{
		creature: creature,
	}
}

func (m *TripleDefenseModifier) Add(next CreatureModifier) {
	m.next = next
}

func (m *TripleDefenseModifier) Apply() {
	m.creature.Defense *= 3

	if m.next != nil {
		m.next.Apply()
	}
}
