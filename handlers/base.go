package handlers

var Base *base

type base struct{}

func (b *base) Initialize() *base {
	Base = &base{}
	return Base
}
