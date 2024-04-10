// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	user "EasyGoZero/app/api/user/internal/handler/user"
	"EasyGoZero/app/api/user/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: user.LoginHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Test},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/user/info",
					Handler: user.GetUserInfoHandler(serverCtx),
				},
			}...,
		),
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}