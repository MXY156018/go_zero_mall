package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"go_zero_mall/ShopVueApi/internal/svc"
	"go_zero_mall/ShopVueApi/internal/types"
	"go_zero_mall/database"
	"go_zero_mall/tool"

	"github.com/tal-tech/go-zero/core/logx"
)

type StoreCombination struct {
	Id            int      `json:"id,omitempty"`
	ProductId     int      `json:"product_id,omitempty"`
	Image         string   `json:"image,omitempty"`
	Images        []string `json:"images,omitempty"`
	Title         string   `json:"title,omitempty"`
	People        int      `json:"people,omitempty"`
	Info          string   `json:"info,omitempty"`
	Price         float64  `json:"price,omitempty"`
	Sales         int      `json:"sales,omitempty"`
	Stock         int      `json:"stock,omitempty"`
	IsPostage     int      `json:"is_postage,omitempty"`
	Postage       float64  `json:"postage,omitempty"`
	Description   string   `json:"description,omitempty"`
	StartTime     int      `json:"start_time,omitempty"`
	StopTime      int      `json:"stop_time,omitempty"`
	EffectiveTime int      `json:"effective_time,omitempty"`
	Browse        int      `json:"browse,omitempty"`
	UnitName      string   `json:"unit_name,omitempty"`
}
type CombinationDetail struct {
	EffectiveTime int    `json:"effective_time"`
	Id            int    `json:"id"`
	Image         string `json:"image"`
	People        int    `json:"people"`
	Price         string `json:"price"`
	ProductPrice  string `json:"product_price"`
	Title         string `json:"title"`
	UnitName      string `json:"unit_name"`
}

type CombinationDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCombinationDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) CombinationListLogic {
	return CombinationListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CombinationListLogic) CombinationDetail(req types.CombinationDatailRequest) (types.Response, error) {
	Detail(req.Id)
	return types.Response{
		Status: 200,
		Msg:    "成功",
		Data:   " GetAll(req.Page, req.Limit)",
	}, nil
}

func Detail(id int) {
	combinationOne := GetCombinationOne(id)

	if !tool.IntToBool(id) || combinationOne.Id == 0 {
		return
	}
	var str []string
	var com StoreCombination
	com = StoreCombination{
		Id:            combinationOne.Id,
		ProductId:     combinationOne.ProductId,
		Image:         combinationOne.Image,
		Title:         combinationOne.Title,
		People:        combinationOne.People,
		Info:          combinationOne.Info,
		Price:         combinationOne.Price,
		Sales:         combinationOne.Sales,
		Stock:         combinationOne.Stock,
		IsPostage:     combinationOne.IsPostage,
		Postage:       combinationOne.Postage,
		Description:   combinationOne.Description,
		StartTime:     combinationOne.StartTime,
		StopTime:      combinationOne.StopTime,
		EffectiveTime: combinationOne.EffectiveTime,
		Browse:        combinationOne.Browse,
		UnitName:      combinationOne.UnitName,
	}
	json.Unmarshal([]byte(combinationOne.Images), &str)

	for i := 0; i < len(str); i++ {
		com.Images = append(com.Images, str[i])
	}
	fmt.Printf("%v\n", com)
}

func GetCombinationOne(id int) types.StoreCombination {
	var conOne types.StoreCombination
	query := "select c.*,s.price as product_price from eb_store_combination c inner join  eb_store_product s on s.id=c.product_id where c.is_show=1 and c.is_del=0 and c.id=?"
	database.DB.Raw(query, id).Find(&conOne)
	return conOne
}
