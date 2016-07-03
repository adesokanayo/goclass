package toy

// Toy is exported type
type Toy struct {
	Name   string
	Weight int
	onHand int
	sold   int
}

//New was exported
func New(name string, weight int) *Toy {
	return &Toy{
		Name:   name,
		Weight: weight,
	}
}

//OnHand is exporetd
func (t *Toy) OnHand() int {
	return t.onHand
}

//copied from answers , not too clear, so I had to read

// UpdateOnHand updates the on hand count and
// returns the current value.
func (t *Toy) UpdateOnHand(count int) int {
	t.onHand += count
	return t.onHand
}

// Sold returns the current number of this
// toy sold.
func (t *Toy) Sold() int {
	return t.sold
}

// UpdateSold updates the sold count and
// returns the current value.
func (t *Toy) UpdateSold(count int) int {
	t.sold += count
	return t.sold
}
