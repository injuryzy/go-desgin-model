package responsibility

import "testing"

func TestName(t *testing.T) {

	handle := BaseHandle{}
	handle1 := BaseHandle1{}

	handle.setNext(&handle1)

	handle.Handle("责任")

}
