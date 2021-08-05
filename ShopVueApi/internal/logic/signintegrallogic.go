package logic

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"go_zero_mall/ShopVueApi/internal/svc"
	"go_zero_mall/ShopVueApi/internal/types"
	"go_zero_mall/database"
	"go_zero_mall/tool"

	"github.com/tal-tech/go-zero/core/logx"
)

type SignIntegralLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignIntegralLogic(ctx context.Context, svcCtx *svc.ServiceContext) SignIntegralLogic {
	return SignIntegralLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SignIntegralLogic) SignIntegral() (types.Response, error) {
	uid := tool.UserInfo.Uid
	signed := GetIsSign(uid, "today")
	if signed {
		return types.Response{
			Status: 200,
			Msg:    "已签到",
			Data:   "",
		}, nil
	}
	integral := Sign(uid)
	if integral != 0 {
		return types.Response{
			Status: 200,
			Msg:    "签到获得" + fmt.Sprint(integral) + "积分",
			Data:   "",
		}, nil
	} else {
		return types.Response{
			Status: 400,
			Msg:    "签到失败",
			Data:   "",
		}, nil
	}
}

//签到
func Sign(uid int) int {
	sign_list := GetSignSystemList()
	if len(sign_list) == 0 {
		return 0
	}
	var user tool.User
	database.DB.Model(&user).Where("uid=?", uid).Find(&user)
	tool.SetUserinfo(user)
	sign_num := 0

	if GetIsSign(uid, "yesterday") {
		if tool.UserInfo.SignNum > len(sign_list)-1 {
			tool.UserInfo.SignNum = 0
		}
	} else {
		tool.UserInfo.SignNum = 0
	}

	for i := 0; i < len(sign_list); i++ {
		if i == tool.UserInfo.SignNum {
			val, _ := strconv.Atoi(sign_list[i].SignNum)
			sign_num = val
			break
		}
	}
	tool.UserInfo.SignNum += 1
	database.DB.Begin()
	res1 := false
	tx := database.DB.Begin()
	if tool.UserInfo.SignNum == len(sign_list) {
		res1 = SetSignData(uid, "连续签到奖励", sign_num, int64(tool.UserInfo.Integral))
	} else {
		res1 = SetSignData(uid, "签到奖励", sign_num, int64(tool.UserInfo.Integral))
	}
	res2 := BcInc(uid, "integral", sign_num, "uid")
	res3 := database.DB.Model(&types.User{}).Where("uid=?", uid).Update("sign_num", tool.UserInfo.SignNum)
	res := res1 && res2 && res3.RowsAffected != 0
	if res {
		tx.Commit()

	} else {
		tx.Rollback()
	}
	if res {
		return sign_num
	} else {
		return 0
	}
}

func SetSignData(uid int, title string, sign_num int, integral int64) bool {
	result := database.DB.Model(&types.User{}).Create(map[string]interface{}{
		"uid":      uid,
		"title":    title,
		"number":   sign_num,
		"balance":  integral,
		"add_time": time.Now().Unix(),
	})
	result2 := Income(title, uid, "intrgral", "sign", sign_num, 0, integral, title)
	if result.RowsAffected > 0 && result2 > 0 {
		return true
	} else {
		return false
	}
}

func Income(title string, uid int, category string, Type string, number int, link_id int, balance int64, mark string) int64 {
	add_time := time.Now().Unix()
	result := database.DB.Model(&types.UserBill{}).Create(map[string]interface{}{
		"title":    title,
		"uid":      uid,
		"link_id":  link_id,
		"categoty": category,
		"type":     Type,
		"number":   number,
		"balance":  balance,
		"mark":     mark,
		"status":   1,
		"pm":       1,
		"add_time": add_time,
	})
	return result.RowsAffected
}

func BcInc(key int, incField string, inc int, keyField string) bool {
	if !tool.IsNumeric(inc) {
		return false
	}
	if keyField == "" {
		keyField = "id"
	}
	intrgral := 0
	intrgral = int(tool.UserInfo.Integral) + inc
	result := database.DB.Model(&types.User{}).Where(keyField+"=?", key).Update(incField, intrgral)
	if result.RowsAffected > 0 {
		return true
	}
	return false
}
