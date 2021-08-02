package logic

import (
	"context"
	"encoding/json"
	"github.com/tal-tech/go-zero/core/logx"
	"go_zero_mall/ShopVueApi/internal/svc"
	"go_zero_mall/ShopVueApi/internal/types"
	"go_zero_mall/database"
	"go_zero_mall/tool"
	"strconv"
	"strings"
	"time"
)

type ShopVueApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewShopVueApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) ShopVueApiLogic {
	return ShopVueApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ShopVueApiLogic) Index(uid int) (types.Response, error) {

	var Data []types.Data
	var result types.Banners1
	tool.GetSystemConfig()
	var logoUrl string
	json.Unmarshal([]byte(tool.SystenConfig["routine_index_logo"]), &logoUrl)
	return types.Response{
		Status: 200,
		Msg:    "ok",
		Data: types.Return{
			Banner:   GetAllValueBanner(Data, result, "routine_home_banner"),
			Menu:     GetAllValueMenu(Data, result, "routine_home_menus"),
			Roll:     GetAllValueRoll(Data, result, "routine_home_roll_news"),
			Activity: GetAllValueActivity(Data, result, "routine_home_activity"),
			Info:     GetAllValueInfo(Data, "routine_index_page", uid),
			Benefit:  GetBeneFit(),
			Lovely: GetBastBanner(Data, "routine_home_new_banner"),
			LikeInfo: GetHotProduct(uid),
			SiteName: tool.SystenConfig["site_name"],
			LogoUrl:  logoUrl,
		},
	}, nil
}

func GetAllValueBanner(Data []types.Data, result types.Banners1, config_name string) []types.Banner {
	row := database.DB.Model(&types.SystemGroupData{}).Select("eb_system_group_data.*,eb_system_group.config_name").Joins("left join eb_system_group on eb_system_group_data.gid=eb_system_group.id").Where("eb_system_group.config_name=?", config_name).Where("eb_system_group_data.status=?", 1).Order("sort desc,id ASC").Scan(&Data)
	if row.RowsAffected == 0 {
		return nil
	}
	var bannerArr []types.Banner
	for _, v := range Data {
		err := json.Unmarshal([]byte(v.Value), &result)

		if err != nil {
			return nil
		}

		var banner1 types.Banner
		banner1 = types.Banner{
			Id: v.Id, Name: result.Name.Value, Pic: result.Pic.Value, Url: result.Url.Value,
		}

		bannerArr = append(bannerArr, banner1)
	}
	return bannerArr
}
func GetAllValueMenu(Data []types.Data, result types.Banners1, config_name string) []types.Menus {
	row := database.DB.Model(&types.SystemGroupData{}).Select("eb_system_group_data.*,eb_system_group.config_name").Joins("left join eb_system_group on eb_system_group_data.gid=eb_system_group.id").Where("eb_system_group.config_name=?", config_name).Where("eb_system_group_data.status=?", 1).Order("sort desc,id ASC").Scan(&Data)
	if row.RowsAffected == 0 {
		return nil
	}
	var menuArr []types.Menus
	for _, v := range Data {
		err := json.Unmarshal([]byte(v.Value), &result)

		if err != nil {
			return nil
		}

		var banner1 types.Menus
		banner1 = types.Menus{
			Id: v.Id, Name: result.Name.Value, Pic: result.Pic.Value, Url: result.Url.Value,
			Show:   result.Show.Value,
			WapUrl: result.WapUrl.Value,
		}

		menuArr = append(menuArr, banner1)
	}
	return menuArr
}

func GetAllValueRoll(Data []types.Data, result types.Banners1, config_name string) []types.Roll {
	row := database.DB.Model(&types.SystemGroupData{}).Select("eb_system_group_data.*,eb_system_group.config_name").Joins("left join eb_system_group on eb_system_group_data.gid=eb_system_group.id").Where("eb_system_group.config_name=?", config_name).Where("eb_system_group_data.status=?", 1).Order("sort desc,id ASC").Scan(&Data)
	if row.RowsAffected == 0 {
		return nil
	}
	var rollArr []types.Roll
	for _, v := range Data {
		err := json.Unmarshal([]byte(v.Value), &result)

		if err != nil {
			return nil
		}

		var banner1 types.Roll
		banner1 = types.Roll{
			Id:     v.Id,
			Url:    result.Url.Value,
			Show:   result.Show.Value,
			WapUrl: result.WapUrl.Value,
			Info:   result.Info.Value,
		}

		rollArr = append(rollArr, banner1)
	}
	return rollArr
}
func GetAllValueActivity(Data []types.Data, result types.Banners1, config_name string) []types.Activity {
	row := database.DB.Model(&types.SystemGroupData{}).Select("eb_system_group_data.*,eb_system_group.config_name").Joins("left join eb_system_group on eb_system_group_data.gid=eb_system_group.id").Where("eb_system_group.config_name=?", config_name).Where("eb_system_group_data.status=?", 1).Order("sort desc,id ASC").Scan(&Data)
	if row.RowsAffected == 0 {
		return nil
	}
	var acticityArr []types.Activity
	for _, v := range Data {
		err := json.Unmarshal([]byte(v.Value), &result)

		if err != nil {
			return nil
		}

		var banner1 types.Activity
		banner1 = types.Activity{
			Id:      v.Id,
			Pic:     result.Pic.Value,
			Title:   result.Title.Value,
			Info:    result.Info.Value,
			Link:    result.Link.Value,
			WapLink: result.WapLink.Value,
		}

		acticityArr = append(acticityArr, banner1)
	}
	return acticityArr
}
func GetAllValueInfo(Data []types.Data, config_name string, uid int) types.Info {
	var result types.Banners2
	var result2 []types.Banners2
	row := database.DB.Model(&types.SystemGroupData{}).Select("eb_system_group_data.*,eb_system_group.config_name").Joins("left join eb_system_group on eb_system_group_data.gid=eb_system_group.id").Where("eb_system_group.config_name=?", config_name).Where("eb_system_group_data.status=?", 1).Order("sort desc,id ASC").Scan(&Data)
	if row.RowsAffected == 0 {
		return types.Info{}
	}

	for _, v := range Data {
		err := json.Unmarshal([]byte(v.Value), &result)
		if err != nil {
			return types.Info{}
		}
		result2 = append(result2, result)
	}
	logoUrl := tool.SystenConfig["routine_index_logo"]
	if !strings.Contains(logoUrl, "http") {
		logoUrl = tool.SystenConfig["site_url"] + logoUrl
	}
	fastNumber := result2[0].FastNumber.Value
	bastNumber := result2[0].BastNumber.Value

	firstNumber := result2[0].FirstNumber.Value

	var info types.Info
	info = types.Info{
		FastInfo:   result2[0].FastInfo.Value,
		BastInfo:   result2[0].BastInfo.Value,
		FirstInfo:  result2[0].FirstInfo.Value,
		SalesInfo:  result2[0].SalesInfo.Value,
		FastList:   GetStoreCateGory(fastNumber),
		BastList:   GetBastProduct(bastNumber, uid),
		FirstList:  GetFirstList(firstNumber, uid),
		BastBanner: GetBastBanner(Data, "routine_home_bast_banner"),

	}

	return info
}
func GetStoreCateGory(fastNumber string) []types.StoreCategory {
	var store_category []types.StoreCategory
	limit, err := strconv.Atoi(fastNumber)
	if err != nil {
		return store_category
	}
	database.DB.Where("pid > ? and is_show = ?", 0, 1).Select("id,cate_name,pid,pic").Order("sort desc").Limit(limit).Find(&store_category)
	return store_category
}
func GetBastProduct(bastNumber string, uid int) []types.StoreProduct {
	var bastList []types.StoreProduct
	limit, err := strconv.Atoi(bastNumber)
	if err != nil {
		return bastList
	}
	database.DB.Model(&bastList).Select("id,image,store_name,cate_id,price,ot_price,IFNULL(sales,0) + IFNULL(ficti,0) as sales,unit_name").Where("is_best = 1 AND is_del = 0 AND mer_id = 0 AND stock > 0 AND is_show = 1").Order("sort desc,id desc").Limit(limit).Find(&bastList)

	levelId := GetUserLevel(uid, 0)

	var discount float64 = 0
	if levelId > 0 {
		discount = GetUserLevelInfo(levelId, "discount")
		discount = 1 - discount/100
	} else {
		var system_user_level types.SystemUserLevel
		database.DB.Model(&system_user_level).Where("id=1 and is_show=1 and is_del=0").Find(&system_user_level)
		discount = system_user_level.Discount
		discount = 1 - discount/100
	}

	for i := 0; i < len(bastList); i++ {
		if bastList[i].Price > 0 {
			bastList[i].VipPrice = bastList[i].Price - bastList[i].Price*discount
		}
	}
	return bastList
}
func GetFirstList(firstNumber string, uid int) []types.Firstlist {
	var bastList types.StoreProduct
	var firstList []types.Firstlist
	limit, err := strconv.Atoi(firstNumber)
	if err != nil {
		return []types.Firstlist{}
	}
	database.DB.Model(&bastList).Select("id,image,store_name,cate_id,price,unit_name,IFNULL(sales,0) + IFNULL(ficti,0) as sales").Where("is_new = 1 AND is_del = 0 AND mer_id = 0 AND stock > 0 AND is_show = 1").Order("sort desc,id desc").Limit(limit).Find(&firstList)
	levelId := GetUserLevel(uid, 0)

	var discount float64 = 0
	if levelId > 0 {
		discount = GetUserLevelInfo(levelId, "discount")
		discount = 1 - discount/100
	} else {
		var system_user_level types.SystemUserLevel
		database.DB.Model(&system_user_level).Where("id=1 and is_show=1 and is_del=0").Find(&system_user_level)
		discount = system_user_level.Discount
		discount = 1 - discount/100
	}

	for i := 0; i < len(firstList); i++ {
		if firstList[i].Price > 0 {
			firstList[i].VipPrice = firstList[i].Price - firstList[i].Price*discount
		}
	}
	return firstList
}
func GetUserLevel(uid int, grade int) int {
	var user_level types.UserLevel
	if grade > 0 {
		database.DB.Model(&user_level).Select("level_id,is_forever,valid_time,id,status,grade").Where("status=1 and is_del=0 and grade < ? and uid=?", grade, uid).Order("grade desc").Find(&user_level)
	} else {
		database.DB.Model(&user_level).Select("level_id,is_forever,valid_time,id,status,grade").Where("status=1 and is_del=0 and uid=?", uid).Order("grade desc").Find(&user_level)
	}

	if user_level.Id == 0 {
		return 0
	}
	if user_level.IsForever == 1 {
		return user_level.Id
	}
	if time.Now().Unix()-user_level.ValidTime > 0 {
		if user_level.Status == 1 {
			database.DB.Model(&user_level).Update("status", 0)
			GetUserLevel(uid, user_level.Grade)
		}
	} else {
		return user_level.Id
	}
	return 0
}
func GetUserLevelInfo(levelId int, keyName string) float64 {
	var user_level types.UserLevel
	var system_user_level types.SystemUserLevel

	database.DB.Model(&user_level).Joins("left join eb_system_user_level on eb_system_user_level.id=eb_user_level.user_id").Where("eb_user_level.id=? and status==1 and is_del==0", levelId).Find(&system_user_level)
	if keyName != "" {
		return system_user_level.Discount
	}
	return 1
}

func GetBastBanner(Data []types.Data, config_name string) []types.BastBanner {
	var result types.BastBanner2
	var banner types.BastBanner
	var banners []types.BastBanner
	row := database.DB.Model(&types.SystemGroupData{}).Select("eb_system_group_data.*,eb_system_group.config_name").Joins("left join eb_system_group on eb_system_group_data.gid=eb_system_group.id").Where("eb_system_group.config_name=?", config_name).Where("eb_system_group_data.status=?", 1).Order("sort desc,id ASC").Scan(&Data)
	if row.RowsAffected == 0 {
		return banners
	}
	for i := 0; i < len(Data); i++ {
		json.Unmarshal([]byte(Data[i].Value), &result)

		banner = types.BastBanner{
			Id:      Data[i].Id,
			Img:     result.Img.Value,
			Comment: result.Comment.Value,
			Link:    result.Link.Value,
			WapLink: result.WapLink.Value,
		}
		banners = append(banners, banner)
	}
	return banners
}

func GetBeneFit() []types.Benefit {
	var benefit []types.Benefit
	database.DB.Model(&types.StoreProduct{}).Select("id,image,store_name,cate_id,price,ot_price,stock,unit_name").Where("is_benefit = 1 AND is_del = 0 AND mer_id = 0 AND stock > 0 AND is_show = 1").Order("sort desc,id desc").Limit(3).Find(&benefit)
	return benefit
}

func GetHotProduct(uid int )[]types.LinkInfo{
	var bastList types.StoreProduct
	var likeinfo []types.LinkInfo

	database.DB.Model(&bastList).Select("id,image,store_name,cate_id,price,unit_name,IFNULL(sales,0) + IFNULL(ficti,0) as sales").Where("is_hot = 1 AND is_del = 0 AND mer_id = 0 AND stock > 0 AND is_show = 1").Order("sort desc,id desc").Limit(3).Find(&likeinfo)
	levelId := GetUserLevel(uid, 0)

	var discount float64 = 0
	if levelId > 0 {
		discount = GetUserLevelInfo(levelId, "discount")
		discount = 1 - discount/100
	} else {
		var system_user_level types.SystemUserLevel
		database.DB.Model(&system_user_level).Where("id=1 and is_show=1 and is_del=0").Find(&system_user_level)
		discount = system_user_level.Discount
		discount = 1 - discount/100
	}

	for i := 0; i < len(likeinfo); i++ {
		if likeinfo[i].Price > 0 {
			likeinfo[i].VipPrice = likeinfo[i].Price - likeinfo[i].Price*discount
		}
	}
	return likeinfo
}