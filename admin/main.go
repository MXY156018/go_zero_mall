package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
	config2 "go_zero_mall/admin/internal/config"
	loginRoutes2 "go_zero_mall/admin/internal/handler/loginRoutes"
	svc2 "go_zero_mall/admin/internal/svc"
	"go_zero_mall/database"
)

var configFile = flag.String("f", "etc/gozeromall-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config2.Config
	conf.MustLoad(*configFile, &c)

	database.DataBase(c.Mysql.SqlDns)

	ctx := svc2.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	loginRoutes2.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
