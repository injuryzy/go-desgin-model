package visitor

type Square struct {
	v Visitor
	A int
	B int
}

func (s *Square) Accept(visitor Visitor) int {
	return visitor.GetSquare(*s)
}
