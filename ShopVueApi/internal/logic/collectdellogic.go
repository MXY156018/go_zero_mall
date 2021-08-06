package logic

import (
	"context"
	"go_zero_mall/ShopVueApi/internal/svc"
	"go_zero_mall/ShopVueApi/internal/types"
	"go_zero_mall/database"
	"go_zero_mall/tool"
	"strings"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
)

type CollectDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCollectDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) CollectDelLogic {
	return CollectDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CollectDelLogic) CollectDel(req types.CollectAddRequest) (types.Response, error) {
	uid := tool.UserInfo.Uid
	if !tool.IntToBool(req.Id) || tool.IsNumeric(req.Id) {
		return types.Response{
			Status: 400,
			Msg:    "参数错误",
			Data:   "",
		}, nil
	}
	if unProductRelation(req.Id, uid, "collect", req.Category) == 1 {
		return types.Response{
			Status: 400,
			Msg:    "产品不存在",
			Data:   "",
		}, nil
	}
	return types.Response{
		Status: 200,
		Msg:    "成功",
		Data:   "",
	}, nil
}

func unProductRelation(productid int, uid int, relationType string, category string) int {
	if !tool.IntToBool(productid) {
		return 1
	}
	relationType = strings.ToLower(relationType)
	category = strings.ToLower(category)
	var data types.StoreProductRelation
	data = types.StoreProductRelation{
		Uid:       uid,
		ProductId: productid,
		Type:      relationType,
		Category:  category,
		AddTime:   int(time.Now().Unix()),
	}

	database.DB.Model(&types.StoreProductRelation{}).Where("uid=? and product_id=? and type=? and category=?", uid, productid, relationType, category).Delete(&data)
	return 2
}
