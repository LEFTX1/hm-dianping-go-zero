package private

import (
	"context"

	"hm-dianping-go-zero/internal/svc"
	"hm-dianping-go-zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LikeBlogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLikeBlogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeBlogLogic {
	return &LikeBlogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LikeBlogLogic) LikeBlog(req *types.LikeBlogRequest) (resp *types.EmptyResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
