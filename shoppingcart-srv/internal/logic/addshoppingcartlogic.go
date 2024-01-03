package logic

import (
	"context"
	"demo/model/mysql"
	"demo/model/redis"
	"demo/shoppingcart-srv/internal/svc"
	"demo/shoppingcart-srv/shoppingCart"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddShoppingCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddShoppingCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddShoppingCartLogic {
	return &AddShoppingCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 添加购物车
func (l *AddShoppingCartLogic) AddShoppingCart(in *shoppingCart.AddShoppingCartRequest) (*shoppingCart.AddShoppingCartResponse, error) {
	// todo: add your logic here and delete this line
	//TODO:添加购物车，逻辑判断购物车中是否已经存在该商品，如果有则提示增加库存，扣减虚拟库存，判断商品虚拟库存是否足够
	stock := redis.LenGoodsList(in.GoodsId)
	if stock > in.GoodsNum {
		return nil, errors.New("商品库存不足，无法添加")
	}
	shop, _ := mysql.GetShoppingCartByGoodsId(in.GoodsId)
	if shop.ID != 0 {
		//商品在购物车中已存在
		shop.GoodsNum += in.GoodsNum
		err := mysql.AddShoppingCartTo(shop)
		if err != nil {
			return nil, errors.New("商品在购物车中已存在，续加库存失败")
		}
	}
	shops := mysql.ShoppingCart{
		UserId:     in.UserId,
		GoodsId:    in.GoodsId,
		GoodsNum:   in.GoodsNum,
		GoodsPrice: in.GoodsPrice,
	}
	err := mysql.AddShoppingCart(&shops)
	if err != nil {
		return nil, errors.New("商品添加购物车失败")
	}
	redis.DeclGoodsList(shops.GoodsId, shops.GoodsNum)
	return &shoppingCart.AddShoppingCartResponse{
		ShopId: int64(shops.ID),
	}, nil
}
