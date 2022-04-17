package handlers

import "github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/config"

var Base *base

type base struct {
	conf *config.Config
}

func (b *base) Initialize(conf *config.Config) {
	Base = &base{
		conf: conf,
	}
}
