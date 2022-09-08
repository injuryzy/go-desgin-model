package strategy

import "testing"

func TestName(t *testing.T) {
	context := Context{}
	for i := 0; i < 2; i++ {
		if i%2 == 0 {
			context.SetStrategy(&Strategy1{})
			context.Dothing()
		} else {
			context.SetStrategy(&Strategy2{})
			context.Dothing()
		}
	}

}
