package logic

import (
	"context"

	"hm-dianping-go-zero/internal/svc"
	"hm-dianping-go-zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QueryVoucherOfShopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQueryVoucherOfShopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryVoucherOfShopLogic {
	return &QueryVoucherOfShopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryVoucherOfShopLogic) QueryVoucherOfShop(req *types.QueryVoucherOfShopRequest) (resp *types.QueryVoucherOfShopReply, err error) {
	// todo: add your logic here and delete this line

	return
}
