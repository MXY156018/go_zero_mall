package handler

import (
	"net/http"

	"go_zero_mall/internal/logic"
	"go_zero_mall/internal/svc"
	"go_zero_mall/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func Go_zero_mallHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGo_zero_mallLogic(r.Context(), ctx)
		resp, err := l.Go_zero_mall(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
