package logic

import (
	"context"
	"go_zero_mall/ShopVueApi/internal/svc"
	"go_zero_mall/ShopVueApi/internal/types"
	"go_zero_mall/database"
	"go_zero_mall/tool"
	"strconv"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
	"gorm.io/gorm"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) UserInfoLogic {

	return UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (types.Response, error) {
	uid := tool.UserInfo.Uid
	k, _ := strconv.Atoi(tool.SystenConfig["store_brokerage_statu"])
	var vip bool
	var vip_id int
	var vip_icon string
	var vip_name string
	if tool.SystenConfig["vip_open"] == "0" {
		vip = false
	} else {
		vipId := GetUserLevel(tool.UserInfo.Uid, 0)
		if vipId > 0 {
			vip = true
		} else {
			vip = false
		}
		vipinfo := GetUserLevelInfo1(vipId)
		if vip {
			vip_id = vipId
			vip_icon = vipinfo.Icon
			vip_name = vipinfo.Name
		}
	}
	recharge := tool.SystenConfig["recharge_switch"]
	rechargeSwitch, _ := strconv.Atoi(recharge) //转换为int
	var switchUserInfoArr []types.User

	if tool.UserInfo.Phone != "" && tool.UserInfo.UserType != "h5" {
		var switchUserInfo types.User
		switchUserInfoArr = append(switchUserInfoArr, types.User(tool.UserInfo))
		result := database.DB.Model(&switchUserInfo).Where("account=?", tool.UserInfo.Phone).Where("user_type='h5'").Find(&switchUserInfo)
		if result.RowsAffected != 0 {
			switchUserInfoArr = append(switchUserInfoArr, switchUserInfo)
		}
	} else if tool.UserInfo.Phone != "" && tool.UserInfo.UserType == "h5" {
		var wechatuserinfo types.User
		result := database.DB.Model(&wechatuserinfo).Where("account=?", tool.UserInfo.Phone).Not("user_type='h5'").Find(&wechatuserinfo)
		if result.RowsAffected != 0 {
			switchUserInfoArr = append(switchUserInfoArr, wechatuserinfo)
		}

	} else if tool.UserInfo.Phone == "" {
		switchUserInfoArr = append(switchUserInfoArr, types.User(tool.UserInfo))
	}
	userinfo := types.UserInfo{
		User:              types.User(tool.UserInfo),
		CouponCount:       GetUserValidCouponCount(),
		Like:              GetUserIdCollect(),
		OrderStatusNum:    GetOrderData(),
		Notice:            GetNotice(),
		Brokerage:         GetBrokerage(uid),
		Recharge:          GetRecharge(uid),
		OrderStatusSum:    GetOrderStatusSum(uid),
		ExtractTotalPrice: UserExtractTotalPrice(uid),
		ExtractPrice:      tool.UserInfo.BrokeragePrice,
		Statu:             k,
		Vip:               vip,
		VipId:             vip_id,
		VipIcon:           vip_icon,
		VipName:           vip_name,
		YesterDay:         int(YesterdayCommissionSum(uid)),
		RechargeSwitch:    rechargeSwitch,
		Adminid:           OrderServiceStatus(uid),
		SwitchUserInfo:    switchUserInfoArr,
	}
	return types.Response{
		Status: 200,
		Msg:    "ok",
		Data:   userinfo,
	}, nil
}
func GetUserValidCouponCount() int64 {
	CheckInvalidCoupon()
	var count int64
	database.DB.Model(&types.StoreCouponUser{}).Where("uid=?", tool.UserInfo.Uid).Where("status=0").Order("is_fail ASC,status ASC,add_time DESC").Count(&count)
	return count
}

func GetUserIdCollect() int64 {
	var count int64
	database.DB.Model(&types.StoreProductRelation{}).Where("uid=?", tool.UserInfo.Uid).Where("type=?", "collect").Count(&count)
	return count
}

func CheckInvalidCoupon() {
	time := time.Now().Unix()
	database.DB.Model(&types.StoreCouponUser{}).Where("end_time<?", time).Where("status = 0").Update("status", 2)
}

//订单信息
func GetOrderData() types.OrderStatusNum {
	uid := tool.UserInfo.Uid
	var data types.OrderStatusNum
	var model = database.DB.Model(&types.StoreOrder{})
	data = types.OrderStatusNum{
		OrderCount:     GetOrderCount(uid),
		SumPrice:       GetSumPrice(uid),
		UnpaidCount:    GetCount(uid, 0, model),
		UnshippedCount: GetCount(uid, 1, model),
		ReceivedCount:  GetCount(uid, 2, model),
		EvaluatedCount: GetCount(uid, 3, model),
		CompleteCount:  GetCount(uid, 4, model),
		RefundCount:    GetCount(uid, -1, model),
	}
	return data
}

//订单支付没有退款  总数
func GetOrderCount(uid int) int64 {
	var count int64
	database.DB.Model(&types.StoreOrder{}).Where("is_del=0 and paid=1 and uid=? ", uid).Where("refund_status=0").Count(&count)
	return count
}

//订单支付没有退款   支付总金额
func GetSumPrice(uid int) float64 {
	var sum float64
	database.DB.Model(&types.StoreOrder{}).Where("is_del=0 and paid=1 and refund_status=0 and uid=?", uid).Pluck("COALESCE(SUM(pay_price), 0) as sum", &sum)
	return sum
}

//获得数量
func GetCount(uid int, status int, model *gorm.DB) int64 {
	var count int64
	StatusByWhere(status, uid, model).Where("is_del=0 and uid=?", uid).Count(&count)
	return count
}

//设置查询条件
func StatusByWhere(status int, uid int, model *gorm.DB) *gorm.DB {
	if model == nil {
		model = database.DB.Model(&types.StoreOrder{})
	}
	IN := []int{1, 2}
	if status == -5 {
		return model
	} else if status == 0 { //未支付
		return model.Where("paid=0 and status=0 and refund_status=0")
	} else if status == 1 {
		return model.Where("paid=1 and status=0 and refund_status=0")
	} else if status == 2 {
		return model.Where("paid=1 and status=1 and refund_status=0")
	} else if status == 3 {
		return model.Where("paid=1 and status=2 and refund_status=0")
	} else if status == 4 {
		return model.Where("paid=1 and status=3 and refund_status=0")
	} else if status == -1 {
		return model.Where("paid=1 and refund_status=1")
	} else if status == -2 {
		return model.Where("paid=1 and refund_status=2")
	} else if status == -3 {
		return model.Where("paid=1").Where("refund_status in ?", IN)
	} else {
		return model
	}
}

//用户通知
func GetNotice() int64 {
	var count int64
	var see_count int64
	uid := tool.UserInfo.Uid
	database.DB.Model(&types.UserNotice{}).Where("uid like ?", "%"+","+string(rune(uid))+","+"%").Where("is_send=1").Count(&count)
	database.DB.Model(&types.UserNoticeSee{}).Where("uid=?", uid).Count(&see_count)
	return count - see_count
}

//获取总佣金
func GetBrokerage(uid int) float64 {
	var sum float64
	database.DB.Model(&types.UserBill{}).Where("uid=?", uid).Where("category='now_money' and type='brokerage' and pm=1 and status=1").Pluck("COALESCE(SUM(number), 0) as sum", &sum)
	return sum
}

//累计充值
func GetRecharge(uid int) float64 {
	var sum float64
	database.DB.Model(&types.UserBill{}).Where("uid=?", uid).Where("category='now_money' and type='recharge' and pm=1 and status=1").Pluck("COALESCE(SUM(number), 0) as sum", &sum)
	return sum
}

//累计消费
func GetOrderStatusSum(uid int) float64 {
	var sum float64
	database.DB.Model(&types.StoreOrder{}).Where("uid=?", uid).Where("is_del=0 and paid=1").Pluck("COALESCE(SUM(pay_price), 0) as sum", &sum)
	return sum
}

//累计提现
func UserExtractTotalPrice(uid int) float64 {
	var sum float64
	database.DB.Model(&types.UserExtract{}).Where("uid=?", uid).Where("status=1").Pluck("COALESCE(SUM(extract_price), 0) as sum", &sum)
	return sum
}

//获取用户VIP信息
func GetUserLevelInfo1(vipId int) types.VipInfo {
	var vip types.UserLevel
	var vipInfo types.VipInfo
	database.DB.Model(&vip).Joins("left join eb_system_user_level on eb_system_user_level.id=eb_user_level.level_id").Where("eb_user_level.id=?", vipId).Where("eb_user_level.status=1 and eb_user_level.is_del=0").Select("eb_system_user_level.id,eb_user_level.add_time,eb_system_user_level.discount,eb_user_level.level_id,eb_system_user_level.name,eb_system_user_level.money,eb_system_user_level.icon,eb_system_user_level.is_pay,eb_system_user_level.grade").Find(&vipInfo)
	return vipInfo
}

//昨天总佣金
func YesterdayCommissionSum(uid int) float64 {
	var sum float64
	a, b := tool.GetLastDay()
	start := tool.StringTransferToTimeStamp(a)
	end := tool.StringTransferToTimeStamp(b)
	database.DB.Model(&types.UserBill{}).Where("category='now_money' and type='brokerage' and pm=1 and status=1").Where("add_time>=? and add_time<=?", start, end).Pluck("COALESCE(SUM(number), 0) as sum", &sum)
	return sum
}

//判断是否是客服
func OrderServiceStatus(uid int) bool {

	var count int64
	database.DB.Model(&types.StoreService{}).Where("uid=?", uid).Where("status=1 and customer=1").Count(&count)
	if count == 0 {
		return false
	} else {
		return true
	}
}
