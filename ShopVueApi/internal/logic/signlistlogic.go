package logic

import (
	"context"

	"go_zero_mall/ShopVueApi/internal/svc"
	"go_zero_mall/ShopVueApi/internal/types"
	"go_zero_mall/database"
	"go_zero_mall/tool"

	"github.com/tal-tech/go-zero/core/logx"
)

type SignListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type userBill struct {
	AddTime string `json:"add_time"`
	Title   string `json:"title"`
	Number  string `json:"number"`
}

func NewSignListLogic(ctx context.Context, svcCtx *svc.ServiceContext) SignUserLogic {
	return SignUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignUserLogic) SignList(req types.CouponsRequest) (types.Response, error) {

	return types.Response{
		Status: 200,
		Msg:    "ok",
		Data:   GetSignList(tool.UserInfo.Uid, req.Page, req.Limit),
	}, nil
}

func GetSignList(uid int, page int, limit int) []userBill {
	var signList []types.UserBill
	database.DB.Model(&signList).Joins("left join eb_user on eb_user.uid=eb_user_bill.uid").Select("eb_user_bill.add_time,eb_user_bill.title,eb_user_bill.number").Where("eb_user_bill.category='integral' and eb_user_bill.type='sign' and eb_user_bill.uid=?", uid).Order("eb_user_bill.add_time desc").Offset((page - 1) * limit).Limit(limit).Find(&signList)
	var bill userBill
	var userBil []userBill
	for i := 0; i < len(signList); i++ {
		bill = userBill{
			AddTime: tool.UnixToString(int64(signList[i].AddTime), "2006-01-02"),
			Title:   signList[i].Title,
			Number:  signList[i].Number,
		}
		userBil = append(userBil, bill)
	}
	return userBil
}
