package logic

import (
	"context"
	"github.com/tal-tech/go-zero/core/logx"
	"go_zero_mall/ShopVueApi/internal/svc"
	"go_zero_mall/ShopVueApi/internal/types"
	"go_zero_mall/database"
)

type StoreCategory struct {
	Id       int    `json:"id"`
	CateName string `json:"cate_name"`
	Pid      int    `json:"pid"`
}
type ProductsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProductsLogic {
	return ProductsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProductsLogic) Products(req types.ProductsRequest) (types.Response, error) {
	var storeproductcate []types.StoreProductCate
	var storeproducts []types.StoreProduct
	var product_ids []int
	model := database.DB.Model(&storeproducts).Where("is_del=0 and is_show=1 and mer_id=0")
	if req.Sid != "0" {
		database.DB.Model(&storeproductcate).Where("cate_id", req.Sid).Select("product_id").Scan(&storeproductcate)
		for i := 0; i < len(storeproductcate); i++ {
			product_ids = append(product_ids, storeproductcate[i].ProductId)
		}
		if len(product_ids) > 0 {
			model = model.Where("id IN ?", product_ids)
		} else {
			model = model.Where("cate_id=-1")
		}
	} else if req.Cid > 0 {
		sidsr := PidBySidList(req.Cid)
		model = model.Where("cate_id in ?", sidsr)
	}
	if req.Keyword != "" {
		model = model.Where("keyword||store_name like ?", "%"+req.Keyword+"%")
	}
	if req.News!=0 {
		model=model.Where("is_new=1")
	}

	baseOrder:=""
	if req.PriceOrder=="desc"{
		baseOrder="price DESC"
	}else if req.PriceOrder=="asc"{
		baseOrder="price ASC"
	}
	if req.SalesOrder=="desc"{
		baseOrder="sales DESC"
	}else if req.SalesOrder=="asc"{
		baseOrder="sales ASC"
	}
	if baseOrder!=""{
		baseOrder = baseOrder+", "
	}
	model=model.Order(baseOrder+"sort desc,add_time desc")
	model.Select("id,store_name,cate_id,image,IFNULL(sales,0) + IFNULL(ficti,0) as sales,price,stock").Limit(req.Limit).Offset((req.Page-1)*req.Limit).Find(&storeproducts)

	if req.Type>0{
		//for i := 0; i < len(storeproducts); i++ {
		//
		//}
	}
	uid:=0
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

	for i := 0; i < len(storeproducts); i++ {
		if storeproducts[i].Price > 0 {
			storeproducts[i].VipPrice = storeproducts[i].Price - storeproducts[i].Price*discount
		}
	}
	return types.Response{
		Status: 200,
		Msg:    "ok",
		Data:   storeproducts,
	}, nil
}

func PidBySidList(cid int) []int {
	var sids []StoreCategory
	database.DB.Model(&sids).Where("pid=?", cid).Select("id,cate_name,pid").Find(&sids)
	var sidsr []int
	if len(sids) > 0 {
		for i := 0; i < len(sids); i++ {
			sidsr = append(sidsr, sids[i].Id)
		}
	}
	return sidsr
}
