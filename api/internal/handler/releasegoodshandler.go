package handler

import (
	"net/http"

	"demo/api/internal/logic"
	"demo/api/internal/svc"
	"demo/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ReleaseGoodsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RelaseGoodsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewReleaseGoodsLogic(r.Context(), svcCtx)
		resp, err := l.ReleaseGoods(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
