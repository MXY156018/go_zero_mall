package tool

import (
	"crypto/md5"
	"encoding/hex"
	"go_zero_mall/database"
	"strconv"
	"strings"

	lacia "github.com/jialanli/lacia/utils"
)

type SystemConfig struct {
	Id          int    `json:"id"`
	MenuName    string `json:"menu_name"`
	Type        string `json:"type"`
	InputType   string `json:"input_type"`
	ConfigTabId int    `json:"config_tab_id"`
	Parameter   string `json:"parameter"`
	UploadType  int    `json:"upload_type"`
	Required    string `json:"required"`
	Width       int    `json:"width"`
	High        int    `json:"high"`
	Value       string `json:"value"`
	Info        string `json:"info"`
	Desc        string `json:"desc"`
	Sort        int    `json:"sort"`
	Status      int    `json:"status"`
}
type User struct {
	Uid            int     `json:"uid"`
	RealName       string  `json:"real_name"`
	Birthday       int     `json:"birthday"`
	CardId         string  `json:"card_id"`
	PartnerId      int     `json:"partner_id"`
	GroupId        int     `json:"group_id"`
	Nickname       string  `json:"nickname"`
	Avatar         string  `json:"avatar"`
	Phone          string  `json:"phone"`
	NowMoney       float64 `json:"now_money"`
	BrokeragePrice float64 `json:"brokerage_price"`
	Integral       float64 `json:"integral"`
	SignNum        int     `json:"sign_num"`
	Level          int     `json:"level"`
	SpreadUid      int     `json:"spread_uid"`
	SpreadTime     int     `json:"spread_time"`
	UserType       string  `json:"user_type"`
	IsPromoter     int     `json:"is_promoter"`
	PayCount       int     `json:"pay_count"`
	SpreadCount    int     `json:"spread_count"`
	Addres         string  `json:"addres"`
	Adminid        int     `json:"adminid"`
	LoginType      string  `json:"login_type"`
	Status         int     `json:"status"`
}

var SystenConfig map[string]string
var UserInfo User

func Md5Transfer(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
func GetSystemConfig() {
	var config []SystemConfig
	result := database.DB.Model(&config).Find(&config)
	if result.RowsAffected == 0 {
		return
	}
	SystenConfig = make(map[string]string)
	for i := 0; i < len(config); i++ {
		if lacia.CalcStrFrequencyWith(config[i].Value, "\\u") > 0 { //unicode编码
			v, _ := zhToUnicode([]byte(config[i].Value))
			SystenConfig[config[i].MenuName] = string(v)
		} else if config[i].Value == "\"\"" {
			config[i].Value = ""

		} else {
			SystenConfig[config[i].MenuName] = lacia.RemoveX(config[i].Value, "\\")
		}
	}

}

func zhToUnicode(raw []byte) ([]byte, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}

func SetUserinfo(user User) {
	UserInfo = user
}

func SetSiteUrl(images string, siteUrl string) string {
	if len(strings.Trim(siteUrl, "")) == 0 {
		siteUrl = SystenConfig["site_url"]
	}
	domainTop := images[0:4]
	if domainTop == "http" {
		return images
	}
	images = strings.Replace(images, "\\", "/", -1)
	return siteUrl + images
}

func IsNumeric(val interface{}) bool {
	switch val.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
	case float32, float64, complex64, complex128:
		return true
	case string:
		str := val.(string)
		if str == "" {
			return false
		}
		// Trim any whitespace
		str = strings.Trim(str, " \\t\\n\\r\\v\\f")
		if str[0] == '-' || str[0] == '+' {
			if len(str) == 1 {
				return false
			}
			str = str[1:]
		}
		// hex
		if len(str) > 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X') {
			for _, h := range str[2:] {
				if !((h >= '0' && h <= '9') || (h >= 'a' && h <= 'f') || (h >= 'A' && h <= 'F')) {
					return false
				}
			}
			return true
		}
		// 0-9,Point,Scientific
		p, s, l := 0, 0, len(str)
		for i, v := range str {
			if v == '.' { // Point
				if p > 0 || s > 0 || i+1 == l {
					return false
				}
				p = i
			} else if v == 'e' || v == 'E' { // Scientific
				if i == 0 || s > 0 || i+1 == l {
					return false
				}
				s = i
			} else if v < '0' || v > '9' {
				return false
			}
		}
		return true
	}

	return false
}

func IntToBool(a int) bool {
	if a > 0 {
		return true
	} else {
		return false
	}
}
