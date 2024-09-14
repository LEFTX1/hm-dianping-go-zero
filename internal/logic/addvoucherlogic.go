package logic

import (
	"context"

	"hm-dianping-go-zero/internal/svc"
	"hm-dianping-go-zero/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddVoucherLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddVoucherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddVoucherLogic {
	return &AddVoucherLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddVoucherLogic) AddVoucher(req *types.AddVoucherRequest) (resp *types.VoucherResult, err error) {
	// todo: add your logic here and delete this line

	return
}
