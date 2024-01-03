package logic

import (
	"context"
	"demo/common"
	"demo/model/redis"
	"demo/user-srv/user"
	"fmt"

	"demo/api/internal/svc"
	"demo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCaptchaLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCaptchaLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCaptchaLogic {
	return &UserCaptchaLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCaptchaLogic) UserCaptcha(req *types.UserCaptchaReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	// TODO: 查询该用户是否已注册
	res, err := l.svcCtx.User.GetUser(l.ctx, &user.GetUserRequest{Mobile: req.Mobile})
	if err != nil {
		return &types.Response{
			Code:    500,
			Message: "查询失败",
		}, err
	}
	if res.UserId == 0 {
		return &types.Response{
			Code:    201,
			Message: "用户未注册",
		}, nil
	}
	// TODO:接收手机号发送验证码
	code := common.GetCode()
	var value []*string
	value = append(value, &l.svcCtx.Config.Sms.AccessKeyID, &l.svcCtx.Config.Sms.AccessKeySecret)
	fmt.Println("11111111111111", value)
	b := common.Sms(value, code)
	if !b {
		return &types.Response{
			Code:    500,
			Message: "验证码发送失败",
		}, nil
	}
	err = redis.AddCaptcha(req.Mobile, code)
	if err != nil {
		return &types.Response{
			Code:    500,
			Message: "验证码发送失败",
		}, nil

	}
	return &types.Response{
		Code:    200,
		Message: "验证码发送成功",
		Data: map[string]interface{}{
			"user_id": res.UserId,
		},
	}, nil
}
