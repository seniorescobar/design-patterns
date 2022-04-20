package jersey

const (
	PortoJerseyType JerseyType = iota
	BenficaJerseyType
)

type (
	Factory struct {
		jerseys map[JerseyType]Jersey
	}

	JerseyType int
)

func NewJerseyFactory() *Factory {
	return &Factory{
		jerseys: make(map[JerseyType]Jersey),
	}
}

func (f *Factory) GetJersey(jt JerseyType) Jersey {
	if j, ok := f.jerseys[jt]; ok {
		return j
	}

	var j Jersey

	switch jt {
	case PortoJerseyType:
		j = new(PortoJersey)
	case BenficaJerseyType:
		j = new(BenficaJersey)
	}

	f.jerseys[jt] = j

	return j
}
