package jersey

type BenficaJersey struct{}

func (j *BenficaJersey) Color() string {
	return "red"
}

func (j *BenficaJersey) String() string {
	return j.Color()
}
