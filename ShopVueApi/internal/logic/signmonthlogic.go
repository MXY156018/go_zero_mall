package logic

import (
	"context"
	"strings"

	"go_zero_mall/ShopVueApi/internal/svc"
	"go_zero_mall/ShopVueApi/internal/types"
	"go_zero_mall/database"
	"go_zero_mall/tool"

	"github.com/tal-tech/go-zero/core/logx"
)

type List struct {
	Time string `json:"time"`
	Ids  string `json:"ids"`
}
type getlist struct {
	AddTime string `json:"add_time"`
	Title   string `json:"title"`
	Number  string `json:"number"`
}
type signList struct {
	Month string    `json:"month"`
	List  []getlist `json:"list"`
}

type SignMonthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignMonthLogic(ctx context.Context, svcCtx *svc.ServiceContext) SignMonthLogic {
	return SignMonthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//签到配置
func (l *SignMonthLogic) SignMonth(req types.CouponsRequest) (types.Response, error) {
	uid := tool.UserInfo.Uid
	return types.Response{
		Status: 200,
		Msg:    "ok",
		Data:   GetSignList1(uid, req.Page, req.Limit),
	}, nil
}

func GetSignList1(uid int, page int, limit int) []signList {
	var list1 []List
	var data []signList
	var item signList
	if limit == 0 {
		return data
	}
	if page > 0 {
		query := "SELECT FROM_UNIXTIME(add_time,'%Y-%m') as time,group_concat(id SEPARATOR ', ') ids FROM `eb_user_bill` WHERE uid=1 AND (category='integral' and type='sign') GROUP BY `time` ORDER BY time desc LIMIT 3"
		database.DB.Raw(query).Scan(&list1)
		// database.DB.Model(&types.UserBill{}).Where("uid=?", uid).Where("category=? and type=?", "intrgral", "sign").Select("FROM_UNIXTIME(add_time,'%Y-%m') as time,group_concat(id SEPARATOR ', ') ids").Group("time").Order("time desc").Offset((page - 1) * limit).Limit(limit).Find(&list1)
	} else {
		query := "SELECT FROM_UNIXTIME(add_time,'%Y-%m') as time,group_concat(id SEPARATOR ', ') ids FROM `eb_user_bill` WHERE uid=1 AND (category='integral' and type='sign') GROUP BY `time` ORDER BY time desc"
		database.DB.Raw(query).Scan(&list1)
		// database.DB.Model(&types.UserBill{}).Where("uid=?", uid).Where("category=? and type=?", "intrgral", "sign").Select("FROM_UNIXTIME(add_time,'%Y-%m') as time,group_concat(id SEPARATOR ', ') ids").Group("time").Order("time desc").Find(&list1)
	}
	// return list1

	for i := 0; i < len(list1); i++ {
		item = signList{
			Month: list1[i].Time,
			List:  GetList(list1[i].Ids),
		}
		data = append(data, item)
	}
	return data
}

func GetList(ids string) []getlist {
	var list []getlist
	var id []string
	id = strings.Split(ids, ", ")
	database.DB.Model(&types.UserBill{}).Where("id in ?", id).Select("FROM_UNIXTIME(add_time,'%Y-%m-%d') as add_time,title,number").Order("add_time desc").Find(&list)
	return list
}
