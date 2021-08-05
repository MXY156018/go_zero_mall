package logic

import (
	"context"

	"go_zero_mall/ShopVueApi/internal/svc"
	"go_zero_mall/ShopVueApi/internal/types"
	"go_zero_mall/database"
	"go_zero_mall/tool"

	"github.com/tal-tech/go-zero/core/logx"
)

type SignUser struct {
	Uid             int     `json:"uid"`
	Nickname        string  `json:"nckname"`
	Avatar          string  `json:"avatar"`
	NowMoney        float64 `json:"now_money"`
	BrokeragePrice  float64 `json:"brokerage_price"`
	Integral        float64 `json:"integral"`
	SignNum         int     `json:"sign_num"`
	IsPromoter      bool    `json:"is_promoter"`
	PayCount        int     `json:"pay_count"`
	SpreadCount     int     `json:"spread_count"`
	Adminid         int     `json:"adminid"`
	LoginType       string  `json:"login_type"`
	SumSginDay      int64   `json:"sum_sgin_day"`
	IsDaySign       bool    `json:"is_day_sign"`
	IsYesterDaySgin bool    `json:"is_YesterDay_sgin"`
}
type SignUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) SignUserLogic {
	return SignUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignUserLogic) SignUser(req types.SignUserRequest) (types.Response, error) {
	uid := tool.UserInfo.Uid
	var signuser SignUser
	if req.Sign != 0 || req.All != 0 {
		signuser = SignUser{
			Uid:             tool.UserInfo.Uid,
			Nickname:        tool.UserInfo.Nickname,
			Adminid:         tool.UserInfo.Adminid,
			Avatar:          tool.UserInfo.Avatar,
			BrokeragePrice:  tool.UserInfo.BrokeragePrice,
			Integral:        tool.UserInfo.Integral,
			LoginType:       tool.UserInfo.LoginType,
			NowMoney:        tool.UserInfo.NowMoney,
			PayCount:        tool.UserInfo.PayCount,
			SpreadCount:     tool.UserInfo.SpreadCount,
			SumSginDay:      GetSignSumDay(uid),
			IsDaySign:       GetIsSign(uid, "today"),
			IsYesterDaySgin: GetIsSign(uid, "yesterday"),
		}
	}
	if !signuser.IsDaySign && !signuser.IsYesterDaySgin {
		signuser.SignNum = 0
	} else {
		signuser.SignNum = tool.UserInfo.SignNum
	}
	if tool.SystenConfig["store_brokerage_statu"] == "2" {
		signuser.IsPromoter = true
	} else {
		signuser.IsPromoter = false
	}
	return types.Response{
		Status: 200,
		Msg:    "ok",
		Data:   signuser,
	}, nil
}

//获取签到总天数
func GetSignSumDay(uid int) int64 {
	var count int64
	database.DB.Model(&types.UserSign{}).Where("uid=?", uid).Count(&count)
	return count
}

//是否签到
func GetIsSign(uid int, date string) bool {
	var count int64
	var start, end string
	if date == "today" {
		start, end = tool.GetToday()
	} else if date == "yesterday" {
		start, end = tool.GetLastDay()
	}
	startstamp := tool.StringTransferToTimeStamp(start)
	endstamp := tool.StringTransferToTimeStamp(end)
	database.DB.Model(&types.UserSign{}).Where("uid=?", uid).Where("add_time>? and add_time<?", startstamp, endstamp).Count(&count)
	if count > 0 {
		return true
	} else {
		return false
	}
}
