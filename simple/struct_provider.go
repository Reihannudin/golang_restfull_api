package simple

type Fii struct {
}

func NewFii() *Fii {
	return &Fii{}
}

type Boo struct {
}

func NewBoo() *Boo {
	return &Boo{}
}

type Naci struct {
}

func NewNaci() *Naci {
	return &Naci{}
}

type Fibonaci struct {
	*Fii
	*Boo
	*Naci
}
