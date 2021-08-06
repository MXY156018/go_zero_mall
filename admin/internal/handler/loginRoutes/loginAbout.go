package loginRoutes

import (
	login2 "go_zero_mall/admin/internal/handler/login"
	svc2 "go_zero_mall/admin/internal/svc"
	"net/http"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc2.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/admin/login",
				Handler: login2.LoginHandler(serverCtx),
			},
		},
	)
}
