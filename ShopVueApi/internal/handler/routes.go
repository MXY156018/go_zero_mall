// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"go_zero_mall/ShopVueApi/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/index",
				Handler: ShopVueApiHandlerIndex(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/search/keyword",
				Handler: ShopVueApiHandlerKeyword(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/products",
				Handler: HandlerProducts(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/category",
				Handler: HandlerCategory(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/coupons",
				Handler: HandlerCoupons(serverCtx),
			},
		},
	)
	//文章
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/article/banner/list",
				Handler: HandlerArticleBannerList(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/article/category/list",
				Handler: HandlerArticleCategoryList(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/article/hot/list",
				Handler: HandlerArticleHotList(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/article/detail",
				Handler: HandlerArticleDetail(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/article/list/:cid",
				Handler: HandlerArticleList(serverCtx),
			},
		},
	)
}
