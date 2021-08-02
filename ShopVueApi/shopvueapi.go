package main

import (
	"flag"
	"fmt"
	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
	"go_zero_mall/ShopVueApi/internal/config"
	"go_zero_mall/ShopVueApi/internal/handler"
	"go_zero_mall/ShopVueApi/internal/svc"
	"go_zero_mall/database"
)

var configFile = flag.String("f", "etc/ShopVueApi-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	database.DataBase(c.Mysql.SqlDns)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()
	//server.Use(func(r http.HandlerFunc) http.HandlerFunc {
	//	return func(w http.ResponseWriter, r *http.Request) {
	//		w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	//		//w.Header().Add("Access-Control-Allow-Headers", "Content-Type,token,id,authori-zation") //header的类型
	//		//w.Header().Add("Access-Control-Request-Headers", "Origin, X-Requested-With, content-Type, Accept, Authorization") //header的类型
	//		//w.Header().Set("content-type", "application/json")             //返回数据格式是json
	//	}
	//})
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)

	server.Start()
}
