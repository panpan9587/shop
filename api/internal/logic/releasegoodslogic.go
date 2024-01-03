package logic

import (
	"context"
	"demo/goods-srv/goods"
	"fmt"

	"demo/api/internal/svc"
	"demo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReleaseGoodsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReleaseGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReleaseGoodsLogic {
	return &ReleaseGoodsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 发布商品
func (l *ReleaseGoodsLogic) ReleaseGoods(req *types.RelaseGoodsReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.Goods.ReleaseGoods(l.ctx, &goods.ReleaseGoodsRequest{
		GoodsTitle:  req.GoodsTitle,
		GoodsPrice:  req.GoodsPrice,
		GoodsStock:  req.GoodsStock,
		GoodsStatus: req.GoodsStatus,
	})
	if err != nil {
		fmt.Println(err)
		return &types.Response{
			Code:    500,
			Message: "商品发布失败",
		}, err
	}
	if !res.Pong {
		return &types.Response{
			Code:    500,
			Message: "商品发布失败,某些数据可能不正确",
		}, nil
	}
	return &types.Response{
		Code:    200,
		Message: "商品发布成功",
	}, nil
}
