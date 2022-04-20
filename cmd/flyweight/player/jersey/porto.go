package jersey

type PortoJersey struct{}

func (j *PortoJersey) Color() string {
	return "blue"
}

func (j *PortoJersey) String() string {
	return j.Color()
}
