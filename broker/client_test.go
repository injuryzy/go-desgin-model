package broker

import "testing"

func TestName(t *testing.T) {
	mediator := NewStaticMediator()
	p := &PassengerTrain{
		mediator,
	}
	f := &FastTrain{m: mediator}
	p.Arrive()
	f.Arrive()
	p.Depart()
}
