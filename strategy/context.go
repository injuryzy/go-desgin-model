package strategy

type Context struct {
	s Strategy
}

func (c *Context) SetStrategy(s Strategy) {
	c.s = s
}

func (c Context) Dothing() {
	c.s.Print()
}
