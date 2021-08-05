package logic

import (
	"context"
	"encoding/json"

	"go_zero_mall/ShopVueApi/internal/svc"
	"go_zero_mall/ShopVueApi/internal/types"
	"go_zero_mall/database"

	"github.com/tal-tech/go-zero/core/logx"
)

type signconfig struct {
	Day     types.Fileds `json:"day"`
	SignNum types.Fileds `json:"sign_num"`
}

type SignConfig struct {
	Id      int    `json:"id"`
	Day     string `json:"day"`
	SignNum string `json:"sign_num"`
}

type SignConfigLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignConfigLogic(ctx context.Context, svcCtx *svc.ServiceContext) SignConfigLogic {
	return SignConfigLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//签到配置
func (l *SignConfigLogic) SignConfig() (types.Response, error) {

	return types.Response{
		Status: 200,
		Msg:    "ok",
		Data:   GetSignSystemList(),
	}, nil
}

func GetSignSystemList() []SignConfig {
	var result signconfig
	var sign []SignConfig
	var signconfig SignConfig
	var Data []types.Data
	row := database.DB.Model(&types.SystemGroupData{}).Select("eb_system_group_data.*,eb_system_group.config_name").Joins("left join eb_system_group on eb_system_group_data.gid=eb_system_group.id").Where("eb_system_group.config_name=?", "sign_day_num").Where("eb_system_group_data.status=?", 1).Order("sort desc,id ASC").Scan(&Data)
	if row.RowsAffected == 0 {
		return sign
	}

	for i := 0; i < len(Data); i++ {
		json.Unmarshal([]byte(Data[i].Value), &result)

		signconfig = SignConfig{
			Id:      Data[i].Id,
			Day:     result.Day.Value,
			SignNum: result.SignNum.Value,
		}
		sign = append(sign, signconfig)
	}
	return sign
}
