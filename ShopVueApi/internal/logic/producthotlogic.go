package logic

import (
	"context"
	"go_zero_mall/ShopVueApi/internal/svc"
	"go_zero_mall/ShopVueApi/internal/types"
	"go_zero_mall/database"

	"github.com/tal-tech/go-zero/core/logx"
)

type ProductHotList struct {
	CateId    string  `json:"cate_id"`
	Id        int     `json:"id"`
	Image     string  `json:"image"`
	Price     float64 `json:"price"`
	StoreName string  `json:"store_name"`
	UnitName  string  `json:"unit_name"`
}
type ProductHotLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductHotLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductHotLogic {
	return ProductHotLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductHotLogic) ProductHot(req types.CouponsRequest) (types.Response, error) {
	var producthostlist []ProductHotList
	if req.Limit == 0 {
		return types.Response{
			Status: 200,
			Msg:    "ok",
			Data:   producthostlist,
		}, nil
	}
	database.DB.Model(&types.StoreProduct{}).Where("is_hot=1 and is_del=0 and mer_id=0 and stock > 0 and is_show=1").Select("id,image,store_name,cate_id,price,unit_name").Order("sort desc,id desc").Offset((req.Page - 1) * req.Limit).Limit(req.Limit).Find(&producthostlist)
	return types.Response{
		Status: 200,
		Msg:    "ok",
		Data:   producthostlist,
	}, nil
}
