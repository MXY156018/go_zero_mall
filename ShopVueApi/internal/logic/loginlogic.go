package logic

import (
	"context"
	"go_zero_mall/ShopVueApi/internal/svc"
	"go_zero_mall/ShopVueApi/internal/types"
	"go_zero_mall/database"
	"go_zero_mall/tool"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/tal-tech/go-zero/core/logx"
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

func (l *LoginLogic) Login(req types.LoginRequest) (types.Response, error) {
	var user tool.User
	result := database.DB.Model(&user).Where("account=? and pwd=?", req.Account, tool.Md5Transfer(req.Password)).Find(&user)
	if result.RowsAffected == 0 {
		return types.Response{
			Status: 400,
			Msg:    "账号或密码错误",
			Data:   nil,
		}, nil
	} else {

		if req.Password == "1234567" {
			return types.Response{
				Status: 400,
				Msg:    "请修改您的初始密码，再尝试登陆",
				Data:   nil,
			}, nil
		}
	}
	if user.Status == 0 {
		return types.Response{
			Status: 400,
			Msg:    "已被禁止，请联系管理员",
			Data:   nil,
		}, nil
	}
	//获取token
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, int64(user.Uid))
	if err != nil {
		return types.Response{
			Status: 400,
			Msg:    "获取token失败，请稍后再试",
			Data:   "",
		}, nil
	}
	ex_time := now + accessExpire
	timeobj := time.Unix(int64(ex_time), 0)
	date := timeobj.Format("2006-01-02 15:04:05")
	tool.SetUserinfo(user)
	return types.Response{
		Status: 200,
		Msg:    "ok",
		Data: types.LoginResponse{
			ExpiresTime: date,
			Token:       jwtToken,
		},
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
