package exchange

type Exchange struct {
	Name string
}

func NewExchange(name string) *Exchange {
	return &Exchange{
		Name: name,
	}
}
