package login

import (
	"net/http"

	"go_zero_mall/internal/logic/login"
	"go_zero_mall/internal/svc"
	Typeslogin "go_zero_mall/internal/types/login"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func LoginHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req Typeslogin.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := login.NewLoginLogic(r.Context(), ctx)
		resp, err := l.Login(req, r)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
