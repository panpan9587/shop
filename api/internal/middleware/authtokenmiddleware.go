package middleware

import (
	"demo/api/internal/config"
	"demo/api/internal/types"
	"demo/common"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type AuthTokenMiddleware struct {
}

func NewAuthTokenMiddleware() *AuthTokenMiddleware {
	return &AuthTokenMiddleware{}
}

func (m *AuthTokenMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		//验证token不能为空
		token := r.Header.Get("token")
		if token == "" {
			data := types.Response{
				Code:    500,
				Message: "token不能为空",
			}
			httpx.OkJson(w, data)
		}
		//验证token正则
		var c config.Config
		conf.MustLoad("D:/go_daima/gao2/month/api/etc/user.yaml", &c)
		if !common.VerifyToken(token, c.AuthToken.AccessSecret) {
			data := types.Response{
				Code:    500,
				Message: "token验证失败",
			}
			httpx.OkJson(w, data)
		}
		// Passthrough to next handler if need
		next(w, r)
	}
}
