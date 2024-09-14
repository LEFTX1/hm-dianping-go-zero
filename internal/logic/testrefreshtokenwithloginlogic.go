package logic

import (
	"context"

	"hm-dianping-go-zero/internal/svc"
	"hm-dianping-go-zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestRefreshTokenWithLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestRefreshTokenWithLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestRefreshTokenWithLoginLogic {
	return &TestRefreshTokenWithLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestRefreshTokenWithLoginLogic) TestRefreshTokenWithLogin() (resp *types.TestResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
