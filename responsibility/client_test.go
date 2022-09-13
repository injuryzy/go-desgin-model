package responsibility

import "testing"

func TestName(t *testing.T) {

	tv
	handle := BaseHandle{}
	handle1 := BaseHandle1{}

	handle.setNext(&handle1)

	handle.Handle("责任")

}
