package private

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hm-dianping-go-zero/internal/logic/private"
	"hm-dianping-go-zero/internal/svc"
	"hm-dianping-go-zero/internal/types"
)

func AddSeckillVoucherHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.AddVoucherRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := private.NewAddSeckillVoucherLogic(r.Context(), svcCtx)
		resp, err := l.AddSeckillVoucher(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
