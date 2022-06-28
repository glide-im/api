package category

import (
	route "github.com/glide-im/api/internal/api/router"
	"github.com/glide-im/api/internal/dao/common"
	"github.com/glide-im/api/internal/dao/wrapper/category"
)

type CategoryApi struct {
}

// 分类列表
func (a *CategoryApi) List(ctx *route.Context) error {
	model := category.CategoryDao.GetModel(1)
	list := []category.Category{}
	model.Order("weight desc").Find(&list)

	ctx.ReturnSuccess(list)
	return nil
}

// 分类新增
func (a *CategoryApi) Store(ctx *route.Context, request *CategoryStoreRequest) error {
	model := category.CategoryDao.GetModel(1)
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
func (a *CategoryApi) Update(ctx *route.Context, request *CategoryStoreRequest) error {
	model := category.CategoryDao.GetModel(1)
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
	model := category.CategoryDao.GetModel(1)
	id := ctx.Context.Param("id")
	model.Where("id = ?", id).Delete(&category.Category{})
	ctx.ReturnSuccess(nil)
	return nil
}

// 分类排序
func (a *CategoryApi) Order(ctx *route.Context, request *CategoryOrderRequest) error {
	orders := request.Orders
	for _, order := range orders {
		model := category.CategoryDao.GetModel(1)
		model.Where("id = ?", order.ID).Update("weight", order.Weight)
	}
	ctx.ReturnSuccess(nil)
	return nil
}
