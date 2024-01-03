package logic

import (
	"context"
	"demo/model/mysql"
	"errors"

	"demo/goods-srv/goods"
	"demo/goods-srv/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGoodsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsInfoLogic {
	return &GetGoodsInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查看商品详情
func (l *GetGoodsInfoLogic) GetGoodsInfo(in *goods.GetGoodsInfoRequest) (*goods.GetGoodsInfoResponse, error) {
	// todo: add your logic here and delete this line
	g, err := mysql.GetGoodsById(in.GoodsId)
	if err != nil {
		return nil, errors.New("数据查询失败")
	}
	if g == nil {
		return nil, errors.New("商品详情查询失败")
	}
	return &goods.GetGoodsInfoResponse{
		Pong: &goods.GoodsInfo{
			GoodsId:     int64(g.ID),
			GoodsTitle:  g.GoodsTitle,
			GoodsPrice:  g.GoodsPrice,
			GoodsStock:  g.GoodsStock,
			GoodsStatus: g.GoodsStatus,
		},
	}, nil
}
