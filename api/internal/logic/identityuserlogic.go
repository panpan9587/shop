package logic

import (
	"context"
	"demo/user-srv/user"

	"demo/api/internal/svc"
	"demo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type IdentityUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewIdentityUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IdentityUserLogic {
	return &IdentityUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *IdentityUserLogic) IdentityUser(req *types.IdentityUserReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.User.IdentityUser(l.ctx, &user.IdentityUserRequest{
		UserId:   req.UserId,
		RealName: req.RealName,
		CardNo:   req.CardNo,
	})
	if err != nil {
		return &types.Response{
			Code:    500,
			Message: "实名认证失败",
		}, err
	}
	if !res.Pong {
		return &types.Response{
			Code:    500,
			Message: "实名认证失败",
		}, nil
	}
	return &types.Response{
		Code:    200,
		Message: "实名认证成功",
	}, err
}
