// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"go_zero_mall/ShopVueApi/internal/svc"

	"github.com/tal-tech/go-zero/rest"
	"strings"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	//这里注册文件服务
	dirlevel := []string{":1", ":2", ":3", ":4", ":5", ":6", ":7", ":8"}
	patern := "/Upload/"
	dirpath := "./Upload/"
	for i := 1; i < len(dirlevel); i++ {
		path := patern + strings.Join(dirlevel[:i], "/")
		//最后生成 /asset
		engine.AddRoute(
			rest.Route{
				Method:  http.MethodGet,
				Path:    path,
				Handler: dirhandler(patern, dirpath),
			})
	}
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
	//user
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/login",
				Handler: HandlerLogin(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/sign/config",
				Handler: HandlerSignConfig(serverCtx),
			},
		},
	)
	//用户相关  需登陆
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/userinfo",
				Handler: HandlerUserInfo(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/logout",
				Handler: HandlerLogout(serverCtx),
			},

			{
				Method:  http.MethodGet,
				Path:    "/api/menu/user",
				Handler: HandlerMenuUser(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/sign/user",
				Handler: HandlerSignUser(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/sign/list",
				Handler: HandlerSignList(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/sign/integral",
				Handler: HandlerSignIntegral(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/sign/month",
				Handler: HandlerSignMonth(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
	//用户类 收藏  需登陆
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/collect/user",
				Handler: HandlerCollectUser(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/collect/add",
				Handler: HandlerCollectAdd(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/collect/del",
				Handler: HandlerCollectDel(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/collect/all",
				Handler: HandlerCollectAll(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
	//产品类  无需登陆
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/product/hot",
				Handler: HandlerProductHot(serverCtx),
			},
		},
	)
	//活动 ----拼团
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/api/combination/list",
				Handler: HandlerCombinationList(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/combination/detail",
				Handler: HandlerCombinationDetail(serverCtx),
			},
		},
	)

}
