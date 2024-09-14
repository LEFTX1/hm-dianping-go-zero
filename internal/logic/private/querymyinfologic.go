package private

import (
	"context"

	"hm-dianping-go-zero/internal/svc"
	"hm-dianping-go-zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMyInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMyInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMyInfoLogic {
	return &QueryMyInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryMyInfoLogic) QueryMyInfo() (resp *types.UserDTOReply, err error) {
	// todo: add your logic here and delete this line

	return
}
