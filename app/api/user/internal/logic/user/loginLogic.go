package user

import (
	"EasyGoZero/app/rpc/user/userRpcModel"
	"context"

	"EasyGoZero/app/api/user/internal/svc"
	"EasyGoZero/app/api/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.ReqLogin) (resp *types.ResLogin, err error) {
	token, err := l.svcCtx.LoginRPC.Login(l.ctx, &userRpcModel.ReqLogin{
		IdentityKind: req.IdentityKind,
		Identifier:   req.Identifier,
		Credential:   req.Credential,
	})
	if err != nil {
		return nil, err
	}
	return &types.ResLogin{Token: token.Token}, nil
}
