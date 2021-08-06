package logic

import (
	"context"
	"go_zero_mall/ShopVueApi/internal/svc"
	"go_zero_mall/ShopVueApi/internal/types"
	"go_zero_mall/database"
	"strings"
	"time"

	"go_zero_mall/tool"

	"github.com/tal-tech/go-zero/core/logx"
	"gorm.io/gorm"
)

type CollectAllLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCollectAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) CollectAllLogic {
	return CollectAllLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CollectAllLogic) CollectAll(req types.CollectAllRequest) (types.Response, error) {
	uid := tool.UserInfo.Uid

	if !tool.IntToBool(len(req.Id)) || tool.IsNumeric(req.Id) {
		return types.Response{
			Status: 400,
			Msg:    "参数错误",
			Data:   "",
		}, nil
	}
	tx := database.DB.Model(&types.StoreProductRelation{}).Begin()
	for i := 0; i < len(req.Id); i++ {
		if ProductRelation1(req.Id[i], uid, "collect", req.Category, tx) == 1 {
			tx.Rollback()
			return types.Response{
				Status: 400,
				Msg:    "产品不存在",
				Data:   "",
			}, nil
		}
	}
	tx.Commit()
	return types.Response{
		Status: 200,
		Msg:    "成功",
		Data:   "",
	}, nil
}

func ProductRelation1(productid int, uid int, relationType string, category string, tx *gorm.DB) int {
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
	var count int64
	database.DB.Model(&types.StoreProductRelation{}).Where("uid=? and product_id=? and type=? and category=?", uid, productid, relationType, category).Count(&count)
	if count > 0 {
		return 2
	}
	tx.Create(&data)
	return 2
}
