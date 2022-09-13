package cmd

type Cmd interface {
	Excute()
}

type OnCmd struct {
	d Driver
}

func (o *OnCmd) Excute() {
	o.d.On()
}

type OffCmd struct {
	d Driver
}

func (o *OffCmd) Excute() {
	o.d.Off()
}
