package logic

import (
	"context"

	"hm-dianping-go-zero/internal/svc"
	"hm-dianping-go-zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestRefreshTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestRefreshTokenLogic {
	return &TestRefreshTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestRefreshTokenLogic) TestRefreshToken() (resp *types.TestResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
