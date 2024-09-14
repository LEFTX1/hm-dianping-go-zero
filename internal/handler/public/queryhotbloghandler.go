package public

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"hm-dianping-go-zero/internal/logic/public"
	"hm-dianping-go-zero/internal/svc"
	"hm-dianping-go-zero/internal/types"
)

func QueryHotBlogHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryHotBlogRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := public.NewQueryHotBlogLogic(r.Context(), svcCtx)
		resp, err := l.QueryHotBlog(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
