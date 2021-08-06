package login

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/tal-tech/go-zero/core/logx"
	svc2 "go_zero_mall/admin/internal/svc"
	Typeslogin2 "go_zero_mall/admin/internal/types/login"
	"go_zero_mall/database"
	"net/http"
	"time"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc2.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc2.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req Typeslogin2.LoginRequest, r *http.Request) (*Typeslogin2.LoginResponse, error) {
	var systemadmin Typeslogin2.SystemAdmin

	row := database.DB.Model(&systemadmin).Where("account = ? and pwd = ?", req.Accout, req.Pwd).Find(&systemadmin)

	if row.RowsAffected == 0 {
		return &Typeslogin2.LoginResponse{

		}, nil
	}

	now := time.Now().Unix()

	accessExpire := l.svcCtx.Config.Auth.AccessExpire

	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, int64(systemadmin.Id))

	if err != nil {
		return nil, err
	}

	fmt.Printf("%v", systemadmin)
	getMenus()
	return &Typeslogin2.LoginResponse{
		UserInfo:     systemadmin,
		Token:        jwtToken,
		Expires_time: accessExpire + now,
		//Menus:
		//Logo :
		//LogoSquare:
		//NewOrderAudisLink:
		//UniqueAuth:
		//Version:
	}, nil
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}


func getMenus(){
	var menus []Typeslogin2.SystemMenus
	row:=database.DB.Model(&menus).Find(&menus)
	if row.RowsAffected==0{

	}
	for i, menu := range menus {
		fmt.Printf("%d%v\n",i,menu)
	}

}
