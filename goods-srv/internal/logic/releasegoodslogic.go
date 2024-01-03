package logic

import (
	"context"
	"demo/goods-srv/goods"
	"demo/goods-srv/internal/svc"
	"demo/model/elasticsearch"
	"demo/model/mysql"
	"demo/model/redis"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReleaseGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewReleaseGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReleaseGoodsLogic {
	return &ReleaseGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 发布商品
func (l *ReleaseGoodsLogic) ReleaseGoods(in *goods.ReleaseGoodsRequest) (*goods.ReleaseGoodsResponse, error) {
	// todo: add your logic here and delete this line
	// 根据业务需求如要设置商品的唯一性，防止重复添加，可以给商品标题添加唯一索引，判断商品是否存在，存在则提示是否增加库存
	g := mysql.Goods{
		GoodsTitle:  in.GoodsTitle,
		GoodsPrice:  in.GoodsPrice,
		GoodsStock:  in.GoodsStock,
		GoodsStatus: in.GoodsStatus,
	}
	err := mysql.AddGoods(&g)
	if err != nil {
		return &goods.ReleaseGoodsResponse{
			Pong: false,
		}, errors.New("商品发布失败")
	}
	redis.AddGoodsList(int64(g.ID), g.GoodsStock)
	//添加es搜素
	elasticsearch.AddEsGoods(&g, g.GoodsTitle)
	return &goods.ReleaseGoodsResponse{
		Pong: true,
	}, nil
}
