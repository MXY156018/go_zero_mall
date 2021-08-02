package login

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/tal-tech/go-zero/core/logx"
	"go_zero_mall/database"
	"go_zero_mall/internal/svc"
	Typeslogin "go_zero_mall/internal/types/login"
	"net/http"
	"time"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req Typeslogin.LoginRequest, r *http.Request) (*Typeslogin.LoginResponse, error) {
	var systemadmin Typeslogin.SystemAdmin

	row := database.DB.Model(&systemadmin).Where("account = ? and pwd = ?", req.Accout, req.Pwd).Find(&systemadmin)

	if row.RowsAffected == 0 {
		return &Typeslogin.LoginResponse{

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
	return &Typeslogin.LoginResponse{
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
	var menus []Typeslogin.SystemMenus
	row:=database.DB.Model(&menus).Find(&menus)
	if row.RowsAffected==0{

	}
	for i, menu := range menus {
		fmt.Printf("%d%v\n",i,menu)
	}

}
