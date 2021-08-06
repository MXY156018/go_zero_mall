package logic

import (
	"context"
	"go_zero_mall/ShopVueApi/internal/svc"
	"go_zero_mall/ShopVueApi/internal/types"
	"go_zero_mall/database"
	"go_zero_mall/tool"

	"github.com/tal-tech/go-zero/core/logx"
)

type CollectList struct {
	Pid       int     `json:"pid"`
	Category  string  `json:"category"`
	StoreName string  `json:"store_name"`
	Price     float64 `json:"price"`
	OtPrice   float64 `json:"ot_price"`
	Sales     int     `json:"sales"`
	Image     string  `json:"image"`
	IsDel     int     `json:"is_del"`
	IsShow    int     `json:"is_show"`
	IsFail    bool    `json:"is_fail"`
}

type CollectUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCollectUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) CollectUserLogic {
	return CollectUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CollectUserLogic) CollectUser(req types.CouponsRequest) (types.Response, error) {
	uid := tool.UserInfo.Uid
	return types.Response{
		Status: 200,
		Msg:    "ok",
		Data:   GetUserCollectProduct(uid, req.Page, req.Limit),
	}, nil
}

func GetUserCollectProduct(uid int, page int, limit int) []CollectList {
	var collectlist []CollectList
	if page > 0 {
		query := "SELECT B.id pid,`A`.`category`,`B`.`store_name`,`B`.`price`,`B`.`ot_price`,`B`.`sales`,`B`.`image`,`B`.`is_del`,`B`.`is_show` FROM `eb_store_product_relation` `A` INNER JOIN `eb_store_product` `B` ON `A`.`product_id`=`B`.`id` WHERE  `A`.`uid` = ?  AND `A`.`type` = 'collect' ORDER BY `A`.`add_time` DESC LIMIT ?,?"
		database.DB.Raw(query, uid, (page-1)*limit, limit).Scan(&collectlist)
	} else {
		query := "SELECT B.id pid,`A`.`category`,`B`.`store_name`,`B`.`price`,`B`.`ot_price`,`B`.`sales`,`B`.`image`,`B`.`is_del`,`B`.`is_show` FROM `eb_store_product_relation` `A` INNER JOIN `eb_store_product` `B` ON `A`.`product_id`=`B`.`id` WHERE  `A`.`uid` = ?  AND `A`.`type` = 'collect' ORDER BY `A`.`add_time` DESC"
		database.DB.Raw(query, uid).Scan(&collectlist)
	}
	for i := 0; i < len(collectlist); i++ {
		if collectlist[i].Pid > 0 {
			if tool.IntToBool(collectlist[i].IsDel) && tool.IntToBool(collectlist[i].IsShow) {
				collectlist[i].IsFail = true
			} else {
				collectlist[i].IsFail = false
			}
		}
	}
	return collectlist
}
