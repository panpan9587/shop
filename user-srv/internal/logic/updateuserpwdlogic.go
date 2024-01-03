package logic

import (
	"context"
	"demo/model/mysql"
	"fmt"
	"github.com/pkg/errors"

	"demo/user-srv/internal/svc"
	"demo/user-srv/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserPwdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserPwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserPwdLogic {
	return &UpdateUserPwdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改用户密码
func (l *UpdateUserPwdLogic) UpdateUserPwd(in *user.UpdateUserPwdRequest) (*user.UpdateUserPwdResponse, error) {
	// todo: add your logic here and delete this line
	err := mysql.UpdateUserPwd(in.NewPassword, in.UserId)
	if err != nil {
		fmt.Println(err)
		return &user.UpdateUserPwdResponse{
			Pong: false,
		}, errors.New("密码更改失败")
	}
	return &user.UpdateUserPwdResponse{
		Pong: true,
	}, nil
}
