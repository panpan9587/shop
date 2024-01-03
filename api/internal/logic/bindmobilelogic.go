package logic

import (
	"context"
	"demo/user-srv/user"

	"demo/api/internal/svc"
	"demo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BindMobileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBindMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindMobileLogic {
	return &BindMobileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 邦定手机号
func (l *BindMobileLogic) BindMobile(req *types.BindMobileReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.User.BindMobile(l.ctx, &user.BindMobileRequest{UserId: req.UserId, Mobile: req.Mobile})
	if err != nil {
		return &types.Response{
			Code:    500,
			Message: "绑定失败",
		}, err
	}
	if !res.Pong {
		return &types.Response{
			Code:    500,
			Message: "绑定失败",
		}, nil
	}
	return &types.Response{
		Code:    200,
		Message: "绑定成功",
	}, nil

}
