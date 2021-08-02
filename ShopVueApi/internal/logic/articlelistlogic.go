package logic

import (
	"context"
	"go_zero_mall/ShopVueApi/internal/svc"
	"go_zero_mall/ShopVueApi/internal/types"
	"go_zero_mall/database"

	"github.com/tal-tech/go-zero/core/logx"
)

type ArticleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ArticleListLogic {
	return ArticleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticleListLogic) List(req types.ArticleListRequets) (types.Response, error) {
	field := "id,title,image_input,visit,from_unixtime(add_time,'%Y-%m-%d %H:%i') as add_time,synopsis,url"
	var bannerlist []types.Article
	database.DB.Model(&bannerlist).Select(field).Where("status=1 and hide=0").Where("CONCAT(',',cid,',')  LIKE '%,$cid,%'").Order("sort desc,add_time desc").Limit(req.Limit).Offset((req.Page - 1) * req.Limit).Find(&bannerlist)

	return types.Response{
		Data: bannerlist,
	}, nil
}
