package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
	"go_zero_mall/database"
	"go_zero_mall/internal/config"
	"go_zero_mall/internal/handler/loginRoutes"
	"go_zero_mall/internal/svc"
)

var configFile = flag.String("f", "etc/gozeromall-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	database.DataBase(c.Mysql.SqlDns)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	loginRoutes.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
