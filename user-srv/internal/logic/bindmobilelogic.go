package logic

import (
	"context"
	"demo/model/mysql"
	"errors"
	"fmt"

	"demo/user-srv/internal/svc"
	"demo/user-srv/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type BindMobileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBindMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindMobileLogic {
	return &BindMobileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 绑定手机号
func (l *BindMobileLogic) BindMobile(in *user.BindMobileRequest) (*user.BindMobileResponse, error) {
	// todo: add your logic here and delete this line

	//TODO:查看手机号是否已经被绑定
	u, _ := mysql.GetUserByMobile(in.Mobile)
	if u == nil {
		return nil, errors.New("手机号信息获取失败")
	}
	if u.ID != 0 {
		return &user.BindMobileResponse{
			Pong: true,
		}, errors.New("该手机已绑定手机号,请勿重复绑定")
	}
	u.Mobile = in.Mobile
	u.ID = uint(in.UserId)
	err := mysql.UpdateUserMobile(u)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("绑定手机号失败")
	}
	if u == nil {
		return nil, errors.New("修改失败")
	}
	return &user.BindMobileResponse{
		Pong: true,
	}, nil
}
