package logic

import (
	"context"
	"demo/goods-srv/goods"

	"demo/api/internal/svc"
	"demo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetGoodsInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsInfoLogic {
	return &GetGoodsInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetGoodsInfoLogic) GetGoodsInfo(req *types.GetGoodsInfoReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.Goods.GetGoodsInfo(l.ctx, &goods.GetGoodsInfoRequest{
		GoodsId: req.GoodsId,
	})
	if err != nil {
		return &types.Response{
			Code:    500,
			Message: "商品详情查询失败",
		}, err
	}
	if res.Pong.GoodsId == 0 {
		return &types.Response{
			Code:    500,
			Message: "商品详情查询失败",
		}, nil
	}
	return &types.Response{
		Code:    200,
		Message: "商品详情",
		Data: map[string]interface{}{
			"GoodsId":     res.Pong.GoodsId,
			"GoodsTitle":  res.Pong.GoodsTitle,
			"GoodsPrice":  res.Pong.GoodsPrice,
			"GoodsStock":  res.Pong.GoodsPrice,
			"GoodsStatus": res.Pong.GoodsStock,
		},
	}, nil

}
