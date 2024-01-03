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

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户是否存在
func (l *GetUserLogic) GetUser(in *user.GetUserRequest) (*user.GetUserResponse, error) {
	// todo: add your logic here and delete this line
	u, err := mysql.GetUserByMobile(in.Mobile)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("查询失败")
	}
	if u == nil {
		return nil, errors.New("未找到该用户")
	}
	return &user.GetUserResponse{
		UserId: int64(u.ID),
	}, nil
}
