package private

import (
	"context"

	"hm-dianping-go-zero/internal/svc"
	"hm-dianping-go-zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddSeckillVoucherLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddSeckillVoucherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddSeckillVoucherLogic {
	return &AddSeckillVoucherLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddSeckillVoucherLogic) AddSeckillVoucher(req *types.AddVoucherRequest) (resp *types.VoucherResult, err error) {
	// todo: add your logic here and delete this line

	return
}
