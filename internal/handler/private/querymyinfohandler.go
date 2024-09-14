package private

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hm-dianping-go-zero/internal/logic/private"
	"hm-dianping-go-zero/internal/svc"
)

func QueryMyInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := private.NewQueryMyInfoLogic(r.Context(), svcCtx)
		resp, err := l.QueryMyInfo()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
