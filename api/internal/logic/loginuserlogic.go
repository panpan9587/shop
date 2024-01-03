package logic

import (
	"context"
	"demo/common"
	"demo/user-srv/user"
	"strconv"
	"time"

	"demo/api/internal/svc"
	"demo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginUserLogic {
	return &LoginUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginUserLogic) LoginUser(req *types.LoginUserReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.User.LoginUser(l.ctx, &user.LoginUserRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return &types.Response{
			Code:    500,
			Message: "登录失败",
		}, err
	}
	if res.UserId == 0 {
		return &types.Response{
			Code:    500,
			Message: "登录失败,用户未注册",
		}, nil
	}
	token, err := common.GetJwtToken(l.svcCtx.Config.AuthToken.AccessSecret, time.Now().Unix(), l.svcCtx.Config.AuthToken.AccessExpire, strconv.FormatInt(res.UserId, 10))
	if err != nil {
		return &types.Response{
			Code:    500,
			Message: "登录失败,请检查网络",
		}, nil
	}
	return &types.Response{
		Code:    200,
		Message: "登录成功",
		Data: map[string]interface{}{
			"token": token,
		},
	}, nil
}
