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

type LoginUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginUserLogic {
	return &LoginUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 登录接口
func (l *LoginUserLogic) LoginUser(in *user.LoginUserRequest) (*user.LoginUserResponse, error) {
	// todo: add your logic here and delete this line
	u, err := mysql.GetUserByUsername(in.Username)
	if err != nil {
		return nil, errors.New("用户查询失败")
	}
	if u == nil {
		return nil, errors.New("用户未注册")
	}
	if u.Password != common.Md5(in.Password) {
		return nil, errors.New("用户密码输入错误")
	}
	return &user.LoginUserResponse{
		UserId: int64(u.ID),
	}, nil
}
