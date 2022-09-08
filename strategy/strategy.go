package strategy

import "fmt"

type Strategy interface {
	Print()
}

type Strategy1 struct {
}

func (s *Strategy1) Print() {
	fmt.Println("s1")
}

type Strategy2 struct {
}

func (s *Strategy2) Print() {
	fmt.Println("s2")
}
