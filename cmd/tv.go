package cmd

import "fmt"

type TV struct {
	IsRuning bool
}

type Driver interface {
	On()
	Off()
}

func (T *TV) On() {
	fmt.Println("on")
}

func (T *TV) Off() {
	fmt.Println("off")
}
