package logic

import (
	"context"
	"go_zero_mall/ShopVueApi/internal/svc"
	"go_zero_mall/ShopVueApi/internal/types"
	"go_zero_mall/database"

	"github.com/tal-tech/go-zero/core/logx"
)

type ArticleBannerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewArticleBannerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) ArticleBannerListLogic {
	return ArticleBannerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ArticleBannerListLogic) BannerList() (types.Response, error) {
	field := "id,title,image_input,visit,from_unixtime(add_time,'%Y-%m-%d %H:%i') as add_time,synopsis,url"
	var bannerlist []types.Article
	database.DB.Model(&bannerlist).Select(field).Where("status=1 and hide=0 and is_banner=1").Order("sort desc,add_time desc").Limit(5).Find(&bannerlist)

	return types.Response{
		Status: 200,
		Msg:    "ok",
		Data:   bannerlist,
	}, nil
}
