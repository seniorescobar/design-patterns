package menu

type (
	Iterator interface {
		HasNext() bool
		Next() Food
	}

	Menu interface {
		CreateIterator() Iterator
	}
)
