package responsibility

import "fmt"

type Handler interface {
	setNext(Handler)
	Handle(r interface{})
}

type BaseHandle struct {
	Next Handler
}

func (b *BaseHandle) setNext(handler Handler) {
	if handler != nil {
		b.Next = handler
	}
}

func (b *BaseHandle) Handle(r interface{}) {
	fmt.Println("baseHandler   ", r)
	if b.Next != nil {
		b.Next.Handle(r)
	}
}

type BaseHandle1 struct {
	Next Handler
}

func (b *BaseHandle1) setNext(handler Handler) {
	if handler != nil {
		b.Next = handler
	}
}

func (b *BaseHandle1) Handle(r interface{}) {
	fmt.Println("baseHandler1   ", r)

	if b.Next != nil {
		b.Next.Handle(r)
	}
}
