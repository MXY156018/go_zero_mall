package logic

import (
	"context"
	"go_zero_mall/ShopVueApi/internal/svc"
	"go_zero_mall/ShopVueApi/internal/types"
	"go_zero_mall/database"

	"github.com/tal-tech/go-zero/core/logx"
)

type CategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) CategoryLogic {
	return CategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryLogic) Category() (types.Response, error) {

	var store_category []types.StoreCategory
	var categoryList []types.StoreCategoryData

	database.DB.Model(&store_category).Where("pid = 0 and is_show = 1").Order("sort desc,id desc").Find(&store_category)
	for i := 0; i < len(store_category); i++ {

		var children []types.StoreCategory
		var item types.StoreCategoryData

		database.DB.Model(&store_category).Where("pid = ?", store_category[i].Id).Find(&children)

		item = types.StoreCategoryData{
			Id:       store_category[i].Id,
			Pid:      store_category[i].Pid,
			CateName: store_category[i].CateName,
			Pic:      store_category[i].Pic,
			Children: children,
		}
		categoryList = append(categoryList, item)
	}
	return types.Response{
		Status: 200,
		Msg:    "ok",
		Data:   categoryList,
	}, nil
}
