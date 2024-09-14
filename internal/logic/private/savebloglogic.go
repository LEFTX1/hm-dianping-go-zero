package private

import (
	"context"

	"hm-dianping-go-zero/internal/svc"
	"hm-dianping-go-zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveBlogLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSaveBlogLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveBlogLogic {
	return &SaveBlogLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveBlogLogic) SaveBlog(req *types.SaveBlogRequest) (resp *types.IdResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
