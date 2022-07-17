package category

import (
	route "github.com/glide-im/api/internal/api/router"
	"github.com/glide-im/api/internal/dao/common"
	"github.com/glide-im/api/internal/dao/wrapper/category"
	"github.com/glide-im/api/internal/pkg/db"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

type CategoryApi struct {
}

// 分类列表
func (a *CategoryApi) List(ctx *route.Context) error {
	model := category.CategoryDao.GetModel(ctx.AppID)
	list := []category.Category{}
	model.Order("weight desc").Find(&list)

	ctx.ReturnSuccess(list)
	return nil
}

// 分类新增
func (a *CategoryApi) Store(ctx *route.Context, request *CategoryStoreRequest) error {
	model := category.CategoryDao.GetModel(ctx.AppID)
	store := category.Category{
		AppID:  1,
		Name:   request.Name,
		Weight: request.Weight,
		Icon:   request.Icon,
	}

	_db := model.Create(store)
	if err := common.JustError(_db); err != nil {
		return err
	}
	ctx.ReturnSuccess(store)
	return nil
}

// 分类更新
func (a *CategoryApi) Updates(ctx *route.Context, request *CategoryUpdateRequest) error {
	categories := request.Categories
	for _, _category := range categories {
		model := category.CategoryDao.GetModel(ctx.AppID)
		store := category.Category{
			AppID:  ctx.AppID,
			Name:   _category.Name,
			Weight: _category.Weight,
			Icon:   _category.Icon,
		}
		var _db *gorm.DB
		if _category.ID == 0 {
			_db = db.DB.Model(&category.Category{}).Create(&store)
		} else {
			_db = model.Where("id = ?", _category.ID).Updates(store)
		}

		if err := common.JustError(_db); err != nil {
			return err
		}
	}

	ctx.ReturnSuccess("")
	return nil
}

// 分类更新
func (a *CategoryApi) Update(ctx *route.Context, request *CategoryStoreRequest) error {
	model := category.CategoryDao.GetModel(ctx.AppID)
	update := category.Category{
		AppID:  1,
		Name:   request.Name,
		Weight: request.Weight,
		Icon:   request.Icon,
	}
	id := ctx.Context.Param("id")
	model.Where("id = ?", id).Updates(update)
	ctx.ReturnSuccess(update)
	return nil
}

// 分类删除
func (a *CategoryApi) Delete(ctx *route.Context) error {
	model := category.CategoryDao.GetModel(ctx.AppID)
	id := ctx.Context.Param("id")
	model.Where("id = ?", id).Delete(&category.Category{})
	ctx.ReturnSuccess(nil)
	return nil
}

// 分类排序
func (a *CategoryApi) Order(ctx *route.Context, request *CategoryOrderRequest) error {
	orders := request.Orders
	for _, order := range orders {
		model := category.CategoryDao.GetModel(ctx.AppID)
		model.Where("id = ?", order.ID).Update("weight", order.Weight)
	}
	ctx.ReturnSuccess(nil)
	return nil
}

// 设置用户的分类
func (a *CategoryApi) SetUserCategory(ctx *route.Context, request *CategoryUserRequest) error {
	categoryIds := request.CategoryIds
	uid := ctx.Context.Param("uid")
	err := category.CategoryUserDao.Updates(cast.ToInt64(uid), categoryIds, ctx.Uid)
	if err != nil {
		return err
	}
	ctx.ReturnSuccess(nil)
	return nil
}
