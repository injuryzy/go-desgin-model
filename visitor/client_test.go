package visitor

import "testing"

func TestName(t *testing.T) {
	visitor := AreaVisitor{}
	square := Square{
		A: 10,
		B: 10,
	}
	accept := square.Accept(&visitor)
	t.Log(accept)
}
