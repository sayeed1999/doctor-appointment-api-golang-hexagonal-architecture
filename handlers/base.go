package handlers

var Base *base

type base struct{}

func (b *base) Initialize() {
	Base = &base{}
}
