package cmd

type Button struct {
	c Cmd
}

func (b *Button) Press() {
	b.c.Excute()
}
