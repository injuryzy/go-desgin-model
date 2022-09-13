package cmd

import "testing"

func TestName(t *testing.T) {
	tv := TV{}
	offButton := Button{}
	onButton := Button{}
	offCmd := OffCmd{
		d: &tv,
	}
	onCmd := OnCmd{
		d: &tv,
	}
	offButton.c = &offCmd
	onButton.c = &onCmd
	offButton.Press()
	onButton.Press()
}
