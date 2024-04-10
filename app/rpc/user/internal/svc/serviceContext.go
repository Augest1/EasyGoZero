package svc

import (
	"EasyGoZero/app/model/userAuthModel"
	"EasyGoZero/app/model/userModel"
	"EasyGoZero/app/rpc/user/internal/config"

	_ "github.com/lib/pq"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config        config.Config
	UserModel     userModel.UserModel
	UserAuthModel userAuthModel.UserAuthModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn("postgres", c.DataSource)
	return &ServiceContext{
		Config:        c,
		UserModel:     userModel.NewUserModel(conn, c.Cache),
		UserAuthModel: userAuthModel.NewUserAuthModel(conn, c.Cache),
	}
}
