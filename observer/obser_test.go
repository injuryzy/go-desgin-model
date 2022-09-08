package observer

import "testing"

func TestName(t *testing.T) {
	publisher := Publisher{}
	sub1 := Sub1{}
	sub2 := Sub2{}
	publisher.AddSubcriber(&sub1)
	publisher.AddSubcriber(&sub2)
	publisher.DelSubcriber(&sub2)
	publisher.Notci()
}
