package observer

import "fmt"

type Subcriber interface {
	Update(v interface{})
}

//订阅者

type Sub1 struct {
}

func (s *Sub1) Update(v interface{}) {
	fmt.Println("sub1", v)
}

type Sub2 struct {
}

func (s *Sub2) Update(v interface{}) {
	fmt.Println("sub2", v)
}
