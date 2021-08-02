package handler

import (
	"net/http"

	"go_zero_mall/ShopVueApi/internal/logic"
	"go_zero_mall/ShopVueApi/internal/svc"
	"go_zero_mall/ShopVueApi/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func HandlerArticleList(ctx *svc.ServiceContext) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ArticleListRequets
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewArticleListLogic(r.Context(), ctx)
		resp, err := l.List(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
