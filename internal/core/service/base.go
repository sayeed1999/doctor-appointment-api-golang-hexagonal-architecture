package service

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/config"
	"github.com/sayeed1999/doctor-appointment-api-golang-hexagonal-architecture/internal/core/domain"
)

type base struct {
	rdb      *redis.Client
	ctx      *context.Context
	mailChan chan domain.MailData
	conf     *config.Config
}

func InitializeBaseService(rdb *redis.Client, ctx *context.Context, mailChan chan domain.MailData, conf *config.Config) *base {
	Base := &base{
		rdb:      rdb,
		ctx:      ctx,
		mailChan: mailChan,
		conf:     conf,
	}
	return Base
}
