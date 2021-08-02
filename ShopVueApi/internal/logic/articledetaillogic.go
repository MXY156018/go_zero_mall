package logic

import (
	"context"
	"go_zero_mall/ShopVueApi/internal/svc"
	"go_zero_mall/ShopVueApi/internal/types"
	"go_zero_mall/database"

	"github.com/tal-tech/go-zero/core/logx"
)

type ArticleDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticleDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) ArticleDetailLogic {
	return ArticleDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticleDetailLogic) Detail(req types.ArticleDetailRequest) (types.Response, error) {
	like := "%, " + req.Id + ",%"
	field := "id,title,image_input,visit,from_unixtime(add_time,'%Y-%m-%d %H:%i') as add_time,synopsis,url"
	var bannerlist []types.Article
	database.DB.Model(&bannerlist).Select(field).Where("status=1 and hide=0").Where("CONCAT(',',cid,',')  LIKE ?", like).Order("sort desc,add_time desc").Offset((req.Page - 1) * req.Limit).Limit(req.Limit).Find(&bannerlist)

	return types.Response{
		Data: bannerlist,
	}, nil
}
