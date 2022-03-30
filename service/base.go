package service

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/domain"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/repository"
)

var Base *base

type base struct {
	repo     repository.IRepository
	rdb      *redis.Client
	ctx      *context.Context
	mailChan chan domain.MailData
}

func (b *base) Initialize(repo repository.IRepository, rdb *redis.Client, ctx *context.Context, mailChan chan domain.MailData) *base {
	Base = &base{
		repo:     repo, // central repository for the whole project
		rdb:      rdb,
		ctx:      ctx,
		mailChan: mailChan,
	}
	return Base
}
