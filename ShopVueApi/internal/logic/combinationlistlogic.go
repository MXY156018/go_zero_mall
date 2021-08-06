package logic

import (
	"context"
	"go_zero_mall/ShopVueApi/internal/svc"
	"go_zero_mall/ShopVueApi/internal/types"
	"go_zero_mall/database"
	"go_zero_mall/tool"
	"time"

	"github.com/tal-tech/go-zero/core/logx"
)

type Combination struct {
	EffectiveTime int    `json:"effective_time"`
	Id            int    `json:"id"`
	Image         string `json:"image"`
	People        int    `json:"people"`
	Price         string `json:"price"`
	ProductPrice  string `json:"product_price"`
	Title         string `json:"title"`
	UnitName      string `json:"unit_name"`
}

type CombinationListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCombinationListLogic(ctx context.Context, svcCtx *svc.ServiceContext) CombinationListLogic {
	return CombinationListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CombinationListLogic) CombinationList(req types.CouponsRequest) (types.Response, error) {

	return types.Response{
		Status: 200,
		Msg:    "成功",
		Data:   GetAll(req.Page, req.Limit),
	}, nil
}

func GetAll(page int, limit int) []Combination {
	var comList []Combination
	nowTime := time.Now().Unix()
	query := "select c.*,s.price product_price from eb_store_combination c inner join eb_store_product s on s.id=c.product_id where c.is_show=1 and c.is_del=0 and c.start_time<? and c.stop_time>? order by c.sort desc,c.id desc limit ?,?"
	database.DB.Raw(query, nowTime, nowTime, (page-1)*limit, limit).Scan(&comList)
	for i := 0; i < len(comList); i++ {
		comList[i].Image = tool.SetSiteUrl(comList[i].Image, "")
	}
	return comList
}
