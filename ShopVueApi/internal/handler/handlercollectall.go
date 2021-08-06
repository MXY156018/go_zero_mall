package handler

import (
	"net/http"

	"go_zero_mall/ShopVueApi/internal/logic"
	"go_zero_mall/ShopVueApi/internal/svc"
	"go_zero_mall/ShopVueApi/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func HandlerCollectAll(ctx *svc.ServiceContext) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CollectAllRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCollectAllLogic(r.Context(), ctx)
		resp, err := l.CollectAll(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
