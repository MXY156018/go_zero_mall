// Code generated by goctl. DO NOT EDIT.
package types

type IndexRequst struct {
	Uid int `json:"uid"`
}

//
type SystemGroupData struct {
	Id    int    `json:"id"`
	Gid   int    `json:"gid"`
	Value string `json:"value"`
	Sort  int    `json:"sort"`
}
type SystemGroup struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Info       string `json:"info"`
	ConfigName string `json:"config_name"`
	Fields     string `json:"fields"`
}

//链表查出的数据
type Data struct {
	Id         int    `json:"id"`
	Gid        int    `json:"gid"`
	Value      string `json:"value"`
	Sort       int    `json:"sort"`
	ConfigName string `json:"config_name"`
}

//item数据结构
type Fileds struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

//banner数据结构
type Banner struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
	Pic  string `json:"pic"`
}

//menu数据结构
type Menus struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Pic    string `json:"pic"`
	Url    string `json:"url"`
	Show   string `json:"show"`
	WapUrl string `json:"wap_url"`
}

//Roll数据结构
type Roll struct {
	Id     int    `json:"id"`
	Info   string `json:"info"`
	Url    string `json:"url"`
	Show   string `json:"show"`
	WapUrl string `json:"wap_url"`
}

//
type Activity struct {
	Id      int    `json:"id"`
	Pic     string `json:"pic"`
	Title   string `json:"title"`
	Info    string `json:"info"`
	Link    string `json:"link"`
	WapLink string `json:"wap_link"`
}

//解析json字符串后的数据结构
type Banners1 struct {
	Id      int    `json:"id"`
	Name    Fileds `json:"name"`
	Pic     Fileds `json:"pic"`
	Url     Fileds `json:"url"`
	Show    Fileds `json:"show"`
	WapUrl  Fileds `json:"wap_url"`
	Info    Fileds `json:"info"`
	Title   Fileds `json:"title"`
	Link    Fileds `json:"link"`
	WapLink Fileds `json:"wap_link"`
}

//返回前端的data的数据结构
type Return struct {
	Banner     []Banner     `json:"banner"`
	Menu       []Menus      `json:"menu"`
	Roll       []Roll       `json:"roll"`
	Info       Info         `json:"info"`
	Activity   []Activity   `json:"activity"`
	Lovely     []BastBanner `json:"lovely"`
	Benefit    []Benefit    `json:"benefit"`
	LikeInfo   []LinkInfo   `json:"likeInfo"`
	LogoUrl    string       `json:"logourl"`
	CouponList string       `json:"couponList"`
	SiteName   string       `json:"site_name"`
}

//返回的结构
type Response struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

type StoreCategory struct {
	Id       int    `json:"id"`
	CateName string `json:"cate_name"`
	Pid      int    `json:"pid"`
	Pic      string `json:"pic"`
}
type Firstlist struct {
	Id        int     `json:"id"`
	Image     string  `json:"image"`
	StoreName string  `json:"store_name"`
	CateId    string  `json:"cate_id"`
	Price     float64 `json:"price"`
	Sales     int     `json:"sales"`
	UnitName  string  `json:"unit_name"`
	VipPrice  float64 `json:"vip_price"`
}

type bastbanner struct {
	Id      int    `json:"id"`
	Img     string `json:"img"`
	Comment string `json:"comment"`
	Link    string `json:"link"`
	WapLink string `json:"wap_link"`
}
type StoreProduct struct {
	Id        int     `json:"id"`
	Image     string  `json:"image"`
	StoreName string  `json:"store_name"`
	CateId    string  `json:"cate_id"`
	Price     float64 `json:"price"`
	OtPrice   float64 `json:"ot_price"`
	Sales     int     `json:"sales"`
	UnitName  string  `json:"unit_name"`
	VipPrice  float64 `json:"vip_price"`
}
type Info struct {
	FastInfo   string          `json:"fastInfo"`
	BastInfo   string          `json:"bastInfo"`
	FirstInfo  string          `json:"firstInfo"`
	SalesInfo  string          `json:"salesInfo"`
	FastList   []StoreCategory `json:"fastList"`
	BastList   []StoreProduct  `json:"bastList"`
	FirstList  []Firstlist     `json:"firstList"`
	BastBanner []BastBanner    `json:"bast_banner"`
}

type Banners2 struct {
	FastInfo    Fileds `json:"fast_info"`
	BastInfo    Fileds `json:"bast_info"`
	FirstInfo   Fileds `json:"first_info"`
	SalesInfo   Fileds `json:"sales_info"`
	FastNumber  Fileds `json:"fast_number"`
	BastNumber  Fileds `json:"bast_number"`
	FirstNumber Fileds `json:"first_number"`
}

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

//会员等级表
type UserLevel struct {
	LevelId   int   `json:"level_id"`
	IsForever int   `json:"is_forever"`
	ValidTime int64 `json:"valid_time"`
	AddTime   int   `json:"add_time"`
	Id        int   `json:"id"`
	Status    int   `json:"status"`
	Grade     int   `json:"grade"`
}
type VipInfo struct {
	Id       int     `json:"id"`
	AddTime  int     `json:"add_time"`
	DisCount float64 `json:"discount"`
	LevelId  int     `json:"level_id"`
	Name     string  `json:"name"`
	Money    float64 `json:"money"`
	Icon     string  `json:"icon"`
	IsPay    int     `json:"is_pay"`
	Grade    int     `json:"grade"`
}
type SystemUserLevel struct {
	Id       int     `json:"id"`
	AddTime  int     `json:"add_time"`
	Discount float64 `json:"discount"`
	LevelId  int     `json:"level_id"`
	Name     string  `json:"name"`
	Money    float64 `json:"money"`
	Icon     string  `json:"icon"`
	IsPay    int     `json:"is_pay"`
	Grade    int     `json:"grade"`
}
type BastBanner struct {
	Id      int    `json:"id"`
	Img     string `json:"img"`
	Comment string `json:"comment"`
	Link    string `json:"link"`
	WapLink string `json:"wap_link"`
}
type BastBanner2 struct {
	Img     Fileds `json:"img"`
	Comment Fileds `json:"comment"`
	Link    Fileds `json:"link"`
	WapLink Fileds `json:"wap_link"`
}

type Benefit struct {
	Id        int     `json:"id"`
	Image     string  `json:"image"`
	StoreName string  `json:"store_name"`
	CateId    string  `json:"cate_id"`
	Price     float64 `json:"price"`
	OtPrice   float64 `json:"ot_price"`
	Stock     int     `json:"stock"`
	UnitName  string  `json:"unit_name"`
}

type LinkInfo struct {
	Id        int     `json:"id"`
	Image     string  `json:"image"`
	StoreName string  `json:"store_name"`
	CateId    string  `json:"cate_id"`
	Price     float64 `json:"price"`
	UnitName  string  `json:"unit_name"`
	VipPrice  float64 `json:"vip_price"`
}

//type StoreCouponIssue struct {
//	ID          int `json:"id"`
//	Cid         int `json:"cid"`
//	StartTime   int `json:"start_time"`
//	EndTime     int `json:"end_time"`
//	TotalCount  int `json:"total_count"`
//	RemainCount int `json:"remain_count"`
//	IsPermanent int `json:"is_permanent"`
//	AddTime     int `json:"add_time"`
//}

type ProductsRequest struct {
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
	Keyword    string `json:"keyword"`
	Sid        string `json:"sid"`
	News       int    `json:"news"`
	PriceOrder string `json:"priceOrder"`
	SalesOrder string `json:"salesOrder"`
	Cid        int    `json:"cid"`
	Type       int    `json:"type"`
}

//商品分类
type StoreProductCate struct {
	Id        int `json:"id"`
	ProductId int `json:"product_id"`
	CateId    int `json:"cate_id"`
}
type StoreProductRelation struct {
	Uid       int    `jsonn:"uid"`
	ProductId int    `json:"product_id"`
	Type      string `json:"type"`
	Category  string `json:"category"`
	AddTime   int    `json:"add_time"`
}
type StoreCategoryData struct {
	Id       int             `json:"id"`
	Pid      int             `json:"pid"`
	CateName string          `json:"cate_name"`
	Pic      string          `json:"pic"`
	Children []StoreCategory `json:"children"`
}

type CouponsRequest struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
type StoreCouponIssue struct {
	Id          int     `json:"id"`
	Cid         int     `json:"cid"`
	Uid         int     `json:"uid"`
	StartTime   int     `json:"start_time"`
	EndTime     int     `json:"end_time"`
	TotalCount  int     `json:"total_count"`
	RemainCount int     `json:"remain_count"`
	Status      int     `json:"status"`
	IsPermanent int     `json:"is_permanent"`
	IsDel       int     `json:"is_del"`
	AddTime     int     `json:"add_time"`
	CouponPrice float64 `json:"coupon_price"`
	UseMinPrice float64 `json:"use_min_price"`
	IsUse       bool    `json:"is_use"`
}
type StoreCouponUser struct {
	Id          int     `json:"id"`
	Cid         int     `json:"cid"`
	Uid         int     `json:"uid"`
	CouponTitle string  `json:"coupon_title"`
	CouponPrice float64 `json:"coupon_price"`
	UseMinPrice float64 `json:"use_min_price"`
	AddTime     int     `json:"add_time"`
	EndTime     int     `json:"end_time"`
	UseTime     int     `json:"use_time"`
	Type        string  `json:"type"`
	Status      int     `json:"status"`
	IsFail      int     `json:"is_fail"`
}
type Article struct {
	Id         int    `json:"id"`
	Title      string `json:"title"`
	ImageInput string `json:"image_input"`
	Visit      string `json:"visit"`
	AddTime    string `json:"add_time"`
	Synopsis   string `json:"synopsis"`
	Url        string `json:"url"`
}
type ArticleCategory struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}
type ArticleListRequets struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
type ArticleDetailRequest struct {
	Id    string `json:"id"`
	Page  int    `json:"page"`
	Limit int    `json:"limit"`
}
type LoginRequest struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}
type LoginResponse struct {
	ExpiresTime string `json:"expires_time"`
	Token       string `json:"token"`
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
type OrderStatusNum struct {
	CompleteCount  int64   `json:"complete_count"`
	EvaluatedCount int64   `json:"evaluated_count"`
	OrderCount     int64   `json:"order_count"`
	ReceivedCount  int64   `json:"received_count"`
	RefundCount    int64   `json:"refund_count"`
	SumPrice       float64 `json:"sum_price"`
	UnpaidCount    int64   `json:"unpaid_count"`
	UnshippedCount int64   `json:"unshipped_count"`
}
type UserInfo struct {
	User              User
	CouponCount       int64          `json:"couponCount"`
	Like              int64          `json:"like"`
	OrderStatusNum    OrderStatusNum `json:"orderStatusNum"`
	Notice            int64          `json:"notice"`
	Brokerage         float64        `json:"brokerage"`
	Recharge          float64        `json:"recharge"`
	OrderStatusSum    float64        `json:"orderStatusSum"`
	ExtractTotalPrice float64        `json:"extractTotalPrice"`
	ExtractPrice      float64        `json:"extractPrice"`
	Statu             int            `json:"statu"`
	Vip               bool           `json:"vip"`
	VipId             int            `json:"vip_id"`
	VipIcon           string         `json:"vip_icon"`
	VipName           string         `json:"vip_name"`
	YesterDay         int            `json:"yester_day"`
	RechargeSwitch    int            `json:"recharge_switch"`
	Adminid           bool           `json:"adminid"`
	Phone             string         `json:"phone"`
	SwitchUserInfo    []User         `json:"switchUserInfo"`
}

type StoreOrder struct {
}

type UserNotice struct {
	Id       int    `json:"id"`
	Uid      string `json:"uid"`
	Type     int    `json:"type"`
	User     string `json:"user"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	IsSend   int    `json:"is_send"`
	SendTime int    `json:"send_time"`
}
type UserNoticeSee struct {
}
type UserBill struct {
	AddTime int64  `json:"add_time"`
	Title   string `json:"title"`
	Number  string `json:"number"`
}
type UserExtract struct {
}
type StoreService struct{}
type SignUserRequest struct {
	Sign     int `json:"sign"`
	Integral int `json:"integral"`
	All      int `json:"all"`
}
type UserSign struct {
}

type CollectAddRequest struct {
	Id       int    `json:"id"`
	Category string `json:"category" gorm:"default:'product'"`
}
type CollectAllRequest struct {
	Id       []int  `json:"id"`
	Category string `json:"category" gorm:"default:'product'"`
}
type CombinationDatailRequest struct {
	Id int `json:"id"`
}
type StoreCombination struct {
	Id            int     `json:"id,omitempty"`
	ProductId     int     `json:"product_id,omitempty"`
	Image         string  `json:"image,omitempty"`
	Images        string  `json:"images,omitempty"`
	Title         string  `json:"title,omitempty"`
	People        int     `json:"people,omitempty"`
	Info          string  `json:"info,omitempty"`
	Price         float64 `json:"price,omitempty"`
	Sales         int     `json:"sales,omitempty"`
	Stock         int     `json:"stock,omitempty"`
	IsPostage     int     `json:"is_postage,omitempty"`
	Postage       float64 `json:"postage,omitempty"`
	Description   string  `json:"description,omitempty"`
	StartTime     int     `json:"start_time,omitempty"`
	StopTime      int     `json:"stop_time,omitempty"`
	EffectiveTime int     `json:"effective_time,omitempty"`
	Browse        int     `json:"browse,omitempty"`
	UnitName      string  `json:"unit_name,omitempty"`
}
