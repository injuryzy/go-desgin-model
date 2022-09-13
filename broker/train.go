package broker

import "fmt"

type Train interface {
	//到达
	Arrive()
	//离开
	Depart()
	// 可以到达
	CanArrive()
}
type PassengerTrain struct {
	m Mediator
}

func (p *PassengerTrain) Arrive() {
	if !p.m.CanArrive(p) {
		fmt.Println("PassengerTrain waitting")
		return
	}
	fmt.Println("PassengerTrain: Arrived")
}

func (p *PassengerTrain) Depart() {
	fmt.Println("PassengerTrain leaving")
	p.m.NotifyAboutDepart()
}

func (p *PassengerTrain) CanArrive() {
	fmt.Println("PassengerTrain can")
	p.Arrive()
}

type FastTrain struct {
	m Mediator
}

func (p *FastTrain) Arrive() {
	if !p.m.CanArrive(p) {
		fmt.Println("FastTrain waitting")
		return
	}
	fmt.Println("FastTrain: Arrived")
}

func (p *FastTrain) Depart() {
	fmt.Println("FastTrain leaving")
	p.m.NotifyAboutDepart()
}

func (p *FastTrain) CanArrive() {
	fmt.Println("FastTrain permitted")
	p.Arrive()
}
