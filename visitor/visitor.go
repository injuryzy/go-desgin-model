package visitor

type Visitor interface {
	GetSquare(Square) int
}

type AreaVisitor struct {
}

func (a *AreaVisitor) GetSquare(square Square) int {
	return square.B * square.A
}
