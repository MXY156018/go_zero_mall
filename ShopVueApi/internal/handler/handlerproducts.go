package handler

import (
	"go_zero_mall/ShopVueApi/internal/types"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"go_zero_mall/ShopVueApi/internal/logic"
	"go_zero_mall/ShopVueApi/internal/svc"
)

func HandlerProducts(ctx *svc.ServiceContext) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ProductsRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewProductsLogic(r.Context(), ctx)
		resp, err := l.Products(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
