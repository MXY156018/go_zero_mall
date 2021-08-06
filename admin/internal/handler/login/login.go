package login

import (
	login2 "go_zero_mall/admin/internal/logic/login"
	svc2 "go_zero_mall/admin/internal/svc"
	Typeslogin2 "go_zero_mall/admin/internal/types/login"
	"net/http"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func LoginHandler(ctx *svc2.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req Typeslogin2.LoginRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := login2.NewLoginLogic(r.Context(), ctx)
		resp, err := l.Login(req, r)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
