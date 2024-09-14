package private

import (
	"context"

	"hm-dianping-go-zero/internal/svc"
	"hm-dianping-go-zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SeckillVoucherLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSeckillVoucherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SeckillVoucherLogic {
	return &SeckillVoucherLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SeckillVoucherLogic) SeckillVoucher(req *types.SeckillVoucherRequest) (resp *types.VoucherResult, err error) {
	// todo: add your logic here and delete this line

	return
}
