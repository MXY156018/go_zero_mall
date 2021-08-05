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

type Menu struct {
	Name   types.Fileds `json:"name"`
	Pic    types.Fileds `json:"pic"`
	Url    types.Fileds `json:"url"`
	WapUrl types.Fileds `json:"wap_url"`
}
type MenuStr struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Pic    string `json:"pic"`
	Url    string `json:"url"`
	WapUrl string `json:"wap_url"`
}

type Return struct {
	RoutineMyMenus []MenuStr `json:"routine_my_menus"`
}

type MenuUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMenuUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) MenuUserLogic {
	return MenuUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MenuUserLogic) MenuUser() (types.Response, error) {
	uid := tool.UserInfo.Uid
	vipOpen := tool.SystenConfig["vip_open"]
	data := GetMenuInfo()

	for i := 0; i < len(data); i++ {
		data[i].Pic = tool.SetSiteUrl(data[i].Pic, "")
		if data[i].Id == 137 && !(tool.SystenConfig["store_brokerage_statu"] == "2" || tool.UserInfo.IsPromoter == 1) {
			data = append(data[:i], data[i+1:]...)
		} else if data[i].Id == 174 && !OrderServiceStatus(uid) {
			data = append(data[:i], data[i+1:]...)
		} else if !OrderServiceStatus(uid) && data[i].WapUrl == "/order/order_cancellation" {
			data = append(data[:i], data[i+1:]...)
		} else if data[i].WapUrl == "/user/vip" && vipOpen == "0" {
			data = append(data[:i], data[i+1:]...)
		} else if data[i].WapUrl == "/customer/index" && !OrderServiceStatus(uid) {
			data = append(data[:i], data[i+1:]...)
		}
	}
	fmt.Printf("%v", data)
	return types.Response{
		Status: 200,
		Msg:    "ok",
		Data: Return{
			RoutineMyMenus: data,
		},
	}, nil
}
func GetMenuInfo() []MenuStr {
	var menu1 Menu
	var menu MenuStr
	var menus []MenuStr
	var Data []types.Data
	row := database.DB.Model(&types.SystemGroupData{}).Select("eb_system_group_data.*,eb_system_group.config_name").Joins("left join eb_system_group on eb_system_group_data.gid=eb_system_group.id").Where("eb_system_group.config_name=?", "routine_my_menus").Where("eb_system_group_data.status=?", 1).Order("sort desc,id ASC").Scan(&Data)
	if row.RowsAffected == 0 {
		return menus
	}
	for i := 0; i < len(Data); i++ {
		json.Unmarshal([]byte(Data[i].Value), &menu1)

		menu = MenuStr{
			Id:     Data[i].Id,
			Name:   menu1.Name.Value,
			Pic:    menu1.Pic.Value,
			Url:    menu1.Url.Value,
			WapUrl: menu1.WapUrl.Value,
		}
		menus = append(menus, menu)
	}
	return menus
}
