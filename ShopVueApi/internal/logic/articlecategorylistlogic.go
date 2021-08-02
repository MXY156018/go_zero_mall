package logic

import (
	"context"
	"go_zero_mall/ShopVueApi/internal/svc"
	"go_zero_mall/ShopVueApi/internal/types"
	"go_zero_mall/database"

	"github.com/tal-tech/go-zero/core/logx"
)

type ArticleCategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticleCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ArticleCategoryListLogic {
	return ArticleCategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticleCategoryListLogic) CategoryList() (types.Response, error) {

	var categorylist []types.ArticleCategory
	database.DB.Model(&categorylist).Select("id,title").Where("hidden=0 and is_del=0 and status=1 and pic=0").Order("sort desc").Find(&categorylist)
	var categorylist1 []types.ArticleCategory
	categorylist1 = append(categorylist1, types.ArticleCategory{
		Id:    0,
		Title: "热门",
	})
	for i := 0; i < len(categorylist); i++ {
		categorylist1 = append(categorylist1, categorylist[i])
	}
	return types.Response{
		Status: 200,
		Msg:    "ok",
		Data:   categorylist1,
	}, nil
}
