package handler

import (
	"go_zero_mall/ShopVueApi/internal/types"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"go_zero_mall/ShopVueApi/internal/logic"
	"go_zero_mall/ShopVueApi/internal/svc"
)

func ShopVueApiHandlerIndex(ctx *svc.ServiceContext) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var req types.IndexRequst
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewShopVueApiLogic(r.Context(), ctx)
		resp, err := l.Index(req.Uid)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
