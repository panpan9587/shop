package handler

import (
	"net/http"

	"demo/api/internal/logic"
	"demo/api/internal/svc"
	"demo/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserCaptchaHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserCaptchaReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserCaptchaLogic(r.Context(), svcCtx)
		resp, err := l.UserCaptcha(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
