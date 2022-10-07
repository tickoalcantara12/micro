package handler

import (
	"github.com/tickoalcantara12/micro/v3/service/api"
	"github.com/tickoalcantara12/micro/v3/service/client"
)

type Context interface {
	Client() client.Client
	Service() *api.Service
	Domain() string
}
