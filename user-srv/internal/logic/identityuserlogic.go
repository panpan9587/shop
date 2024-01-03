package logic

import (
	"context"
	"demo/common"
	"demo/model/mysql"
	"errors"

	"demo/user-srv/internal/svc"
	"demo/user-srv/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type IdentityUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewIdentityUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *IdentityUserLogic {
	return &IdentityUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 实名认证
func (l *IdentityUserLogic) IdentityUser(in *user.IdentityUserRequest) (*user.IdentityUserResponse, error) {
	// todo: add your logic here and delete this line
	//TODO:根据用户查询用户是否已经完成实名认证，如以实名则返回
	ident, _ := mysql.GetUserIdentityByUserId(in.UserId)
	if ident.ID != 0 {
		return nil, errors.New("该用户已完成实名")
	}
	isOk := common.Identity(in.RealName, in.CardNo)
	if !isOk {
		return &user.IdentityUserResponse{
			Pong: false,
		}, errors.New("实名失败")
	}
	idents := &mysql.Identity{
		UserId:   in.UserId,
		RealName: in.RealName,
		CardNo:   in.CardNo,
	}
	err := mysql.AddIdentity(idents)
	if err != nil {
		return &user.IdentityUserResponse{
			Pong: false,
		}, errors.New("实名失败")
	}
	return &user.IdentityUserResponse{
		Pong: true,
	}, nil
}
