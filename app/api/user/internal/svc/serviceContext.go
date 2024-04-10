package svc

import (
	"EasyGoZero/app/api/user/internal/config"
	"EasyGoZero/app/rpc/user/client/login"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	LoginRPC login.Login
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		LoginRPC: login.NewLogin(zrpc.MustNewClient(c.UserRPC)),
	}
}
