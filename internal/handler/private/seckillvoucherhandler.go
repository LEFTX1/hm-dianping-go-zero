package private

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hm-dianping-go-zero/internal/logic/private"
	"hm-dianping-go-zero/internal/svc"
	"hm-dianping-go-zero/internal/types"
)

func SeckillVoucherHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SeckillVoucherRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := private.NewSeckillVoucherLogic(r.Context(), svcCtx)
		resp, err := l.SeckillVoucher(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
