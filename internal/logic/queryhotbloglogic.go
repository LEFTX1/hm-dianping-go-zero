package logic

import (
	"context"

	"hm-dianping-go-zero/internal/svc"
	"hm-dianping-go-zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryHotBlogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryHotBlogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryHotBlogLogic {
	return &QueryHotBlogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryHotBlogLogic) QueryHotBlog(req *types.QueryHotBlogRequest) (resp *types.BlogResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
