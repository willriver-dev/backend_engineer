package builder

type iglooBuilder struct {
	windowType string
	doorType string
	floor int
}

func newIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}

func (b *iglooBuilder) setWindowType()  {
	b.windowType = "Snow"
}

func (b *iglooBuilder) setDoorType()  {
	b.windowType = "Snow"
}

func (b *iglooBuilder) setNumbersFloor()  {
	b.floor = 1
}

func (b *iglooBuilder) getHouse() House {
	return House{
		doorType: b.doorType,
		windowType: b.windowType,
		floor: b.floor,
	}
}