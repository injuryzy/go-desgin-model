package observer

//发布者

type Publisher struct {
	sub []Subcriber
}

//添加订阅者
func (p *Publisher) AddSubcriber(s Subcriber) {
	p.sub = append(p.sub, s)
}

func (p *Publisher) Notci() {
	for _, subcriber := range p.sub {
		subcriber.Update("100")
	}
}

func (p *Publisher) DelSubcriber(s Subcriber) {
	for i, subcriber := range p.sub {
		if subcriber == s {
			temp := p.sub[0:i]
			p.sub = append(temp, p.sub[i+1:]...)
		}
	}
}
