package logic

import (
	"context"

	"go_zero_mall/ShopVueApi/internal/svc"
	"go_zero_mall/ShopVueApi/internal/types"
	"go_zero_mall/database"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
	"gorm.io/gorm"
)

type CouponsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCouponsLogic(ctx context.Context, svcCtx *svc.ServiceContext) CouponsLogic {
	return CouponsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CouponsLogic) Coupons(req types.CouponsRequest) (types.Response, error) {
	uid := 0
	var store_coupon []types.StoreCouponIssue

	db := database.DB.Model(&store_coupon).Scopes(validWhere).Joins("left join eb_store_coupon on eb_store_coupon_issue.cid=eb_store_coupon.id").Select("eb_store_coupon_issue.*,eb_store_coupon.coupon_price,eb_store_coupon.use_min_price").Order("eb_store_coupon.sort desc,eb_store_coupon_issue.id desc").Offset((req.Page - 1) * req.Limit).Limit(req.Limit)

	if uid > 0 {
		db.Joins("left join eb_store_coupon_issue_user on eb_store_coupon_issue.cid=eb_store_coupon_issue_user.issue_coupon_id").Select("uid").Find(&store_coupon)

	} else {
		db.Find(&store_coupon)
	}

	// for i := 1; i < len(store_coupon); i++ {
	// 	println(store_coupon[i].Uid)
	// 	if store_coupon[i].Uid == uid {
	// 		store_coupon[i].IsUse = true
	// 	}

	// }
	return types.Response{
		Status: 200,
		Msg:    "ok",
		Data:   store_coupon,
	}, nil
}

func validWhere(db *gorm.DB) *gorm.DB {
	newTime := time.Now().Unix()
	return db.Where("eb_store_coupon.status=1 and eb_store_coupon.is_del=0").Where("remain_count>0 or is_permanent=1").Where("start_time<? and end_time>?", newTime, newTime).Or("start_time=0 and end_time=0")
}
