package logic

import (
	"context"
	"encoding/json"
	"github.com/tal-tech/go-zero/core/logx"
	"go_zero_mall/ShopVueApi/internal/svc"
	"go_zero_mall/ShopVueApi/internal/types"
	"go_zero_mall/database"
)

type HotSearch struct {
	Title types.Fileds
}
type KeywordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewKeywordLogic(ctx context.Context, svcCtx *svc.ServiceContext) KeywordLogic {
	return KeywordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShopVueApiLogic) Keyword() (types.Response, error) {
	var Data []types.Data
	row := database.DB.Model(&types.SystemGroupData{}).Select("eb_system_group_data.*,eb_system_group.config_name").Joins("left join eb_system_group on eb_system_group_data.gid=eb_system_group.id").Where("eb_system_group.config_name=?", "routine_hot_search").Where("eb_system_group_data.status=?", 1).Order("sort desc,id ASC").Scan(&Data)
	if row.RowsAffected == 0 {
		return types.Response{
			Status: 200,
			Msg:    "ok",
			Data:   "",
		}, nil
	}
	var HotArr []string
	var result HotSearch
	for _, v := range Data {
		err := json.Unmarshal([]byte(v.Value), &result)

		if err != nil {
			return types.Response{
				Status: 200,
				Msg:    "ok",
				Data:   "",
			}, nil
		}
		str := result.Title.Value

		HotArr = append(HotArr, str)
	}

	return types.Response{
		Status: 200,
		Msg:    "ok",
		Data:   HotArr,
	}, nil
}
