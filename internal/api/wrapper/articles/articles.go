package articles

import (
	route "github.com/glide-im/api/internal/api/router"
	"github.com/glide-im/api/internal/dao/common"
	"github.com/glide-im/api/internal/dao/wrapper/articles"
)

type ArticleApi struct {
}

// 文章新增
func (a *ArticleApi) List(ctx *route.Context) error {
	model := articles.ArticleDao.GetModel(1, ctx.Uid)
	articlesList := []articles.Article{}
	model.Order("weight desc").Find(&articlesList)

	ctx.ReturnSuccess(articlesList)
	return nil
}

// 文章新增
func (a *ArticleApi) Store(ctx *route.Context, request *ArticlesStoreRequest) error {
	model := articles.ArticleDao.GetModel(1, ctx.Uid)
	articleStore := articles.Article{
		AppID:     1,
		Uid:       ctx.Uid,
		Title:     request.Title,
		Content:   request.Content,
		PublishAt: request.PublishAT,
		Weight:    request.Weight,
	}

	_db := model.Create(articleStore)
	if err := common.JustError(_db); err != nil {
		return err
	}
	ctx.ReturnSuccess(articleStore)
	return nil
}

// 文章更新
func (a *ArticleApi) Update(ctx *route.Context, request *ArticlesStoreRequest) error {
	model := articles.ArticleDao.GetModel(1, ctx.Uid)
	articleUpdate := articles.Article{
		AppID:     1,
		Uid:       ctx.Uid,
		Title:     request.Title,
		Content:   request.Title,
		PublishAt: request.PublishAT,
		Weight:    request.Weight,
	}
	id := ctx.Context.Param("id")
	model.Where("id = ?", id).Updates(articleUpdate)
	ctx.ReturnSuccess(articleUpdate)
	return nil
}

// 文章删除
func (a *ArticleApi) Delete(ctx *route.Context) error {
	model := articles.ArticleDao.GetModel(1, ctx.Uid)
	id := ctx.Context.Param("id")
	model.Where("id = ?", id).Delete(&articles.Article{})
	ctx.ReturnSuccess(nil)
	return nil
}

// 文章排序
func (a *ArticleApi) Order(ctx *route.Context, request *ArticlesOrderRequest) error {
	orders := request.Orders
	for _, order := range orders {
		model := articles.ArticleDao.GetModel(1, ctx.Uid)
		model.Where("id = ?", order.ID).Update("weight", order.Weight)
	}
	ctx.ReturnSuccess(nil)
	return nil
}
