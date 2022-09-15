package decorator

type IPrize interface {
	GetPrize() int
}

type BasePrice struct {
}

func (b *BasePrice) GetPrize() int {

	return 10
}

type VigePrize struct {
	Price IPrize
}

func (v *VigePrize) GetPrize() int {

	return (v.Price.GetPrize()) + 10
}
