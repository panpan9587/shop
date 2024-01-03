package logic

import (
	"context"
	"demo/api/internal/svc"
	"demo/api/internal/types"
	"demo/common"
	"demo/model/redis"
	"demo/user-srv/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ForgetPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewForgetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ForgetPasswordLogic {
	return &ForgetPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ForgetPasswordLogic) ForgetPassword(req *types.ForgetPasswordReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	//TODO: 接收验证码，进行判断是否正确，进行更改密码操作
	code := redis.GetCaptcha(req.Mobile)
	if code != req.Captcha {
		return &types.Response{
			Code:    500,
			Message: "验证码输入有误，请重新输入",
		}, nil
	}
	res, err := l.svcCtx.User.UpdateUserPwd(l.ctx, &user.UpdateUserPwdRequest{UserId: req.UserId, NewPassword: common.Md5(req.NewPassword)})
	if err != nil {
		return &types.Response{
			Code:    500,
			Message: "密码更改失败",
		}, err
	}
	if !res.Pong {
		return &types.Response{
			Code:    500,
			Message: "密码更改失败",
		}, nil
	}
	return &types.Response{
		Code:    200,
		Message: "密码更改成功",
	}, nil
}
