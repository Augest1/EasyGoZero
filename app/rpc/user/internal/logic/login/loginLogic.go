package loginlogic

import (
	"EasyGoZero/app/rpc/user/internal/svc"
	"EasyGoZero/app/rpc/user/userRpcModel"
	"EasyGoZero/common/middleware/jwt"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *userRpcModel.ReqLogin) (*userRpcModel.ResLogin, error) {
	userAuth, err := l.svcCtx.UserAuthModel.FindOneByIdentityKindIdentifier(l.ctx, in.IdentityKind, in.Identifier)
	if errors.Is(err, sqlx.ErrNotFound) {
		return nil, errors.New("用户不存在")
	}
	if err != nil {
		return nil, err
	}
	if userAuth.Credential != in.Credential {
		return nil, errors.New("credential error")
	}
	newJwt := jwt.NewJwt(l.svcCtx.Config.JWT.AccessSecret, l.svcCtx.Config.JWT.AccessExpire)
	token, err := newJwt.GenJwtToken(&jwt.Payload{
		UserId: userAuth.UserId,
	})
	return &userRpcModel.ResLogin{
		Token: token,
	}, nil
}
