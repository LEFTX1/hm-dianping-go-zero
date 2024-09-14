package private

import (
	"context"

	"hm-dianping-go-zero/internal/svc"
	"hm-dianping-go-zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryMyBlogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryMyBlogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryMyBlogLogic {
	return &QueryMyBlogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryMyBlogLogic) QueryMyBlog(req *types.QueryMyBlogRequest) (resp *types.BlogResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
