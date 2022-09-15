package decorator

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	price := BasePrice{}
	prize := VigePrize{Price: &price}
	fmt.Println(prize.GetPrize())
}
