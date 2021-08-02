package loginRoutes

import (
	"net/http"

	"go_zero_mall/internal/svc"

	"go_zero_mall/internal/handler/login"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/admin/login",
				Handler: login.LoginHandler(serverCtx),
			},
		},
	)
}
