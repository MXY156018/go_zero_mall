package handler

import (

	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
	"go_zero_mall/ShopVueApi/internal/logic"
	"go_zero_mall/ShopVueApi/internal/svc"
)

func ShopVueApiHandlerKeyword(ctx *svc.ServiceContext) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var req struct{}
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewShopVueApiLogic(r.Context(), ctx)
		resp, err := l.Keyword()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
