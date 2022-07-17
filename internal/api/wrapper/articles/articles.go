package articles

import (
	"errors"
	route "github.com/glide-im/api/internal/api/router"
	"github.com/glide-im/api/internal/dao"
	"github.com/glide-im/api/internal/dao/common"
	"github.com/glide-im/api/internal/dao/wrapper/articles"
	"github.com/spf13/cast"
	"time"
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
	model := articles.ArticleDao.GetModel(ctx.AppID, ctx.Uid)
	var at dao.JSONTime
	_t, err := time.Parse(dao.TIME_FORMAT, request.PublishAT)
	if err != nil {
		return errors.New("时间格式不正确")
	}
	at.Time = _t
	articleStore := articles.Article{
		AppID:     ctx.AppID,
		Uid:       ctx.Uid,
		Title:     request.Title,
		Content:   request.Content,
		PublishAt: at,
		Weight:    request.Weight,
	}

	_db := model.Create(&articleStore)
	if err := common.JustError(_db); err != nil {
		return err
	}
	ctx.ReturnSuccess(articleStore)
	return nil
}

// 文章更新
func (a *ArticleApi) Update(ctx *route.Context, request *ArticlesStoreRequest) error {
	model := articles.ArticleDao.GetModel(ctx.AppID, ctx.Uid)

	var at dao.JSONTime
	_t, err := time.Parse(dao.TIME_FORMAT, request.PublishAT)
	if err != nil {
		return errors.New("时间格式不正确")
	}
	at.Time = _t

	articleUpdate := articles.Article{
		AppID:     ctx.AppID,
		Uid:       ctx.Uid,
		Title:     request.Title,
		Content:   request.Content,
		PublishAt: at,
		Weight:    request.Weight,
	}
	id := ctx.Context.Param("id")
	model.Where("id = ?", id).Updates(&articleUpdate)
	articleUpdate.Id = cast.ToInt64(id)

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

// 文章详情
func (a *ArticleApi) Show(ctx *route.Context) error {
	model := articles.ArticleDao.GetModel(ctx.AppID, ctx.Uid)
	id := ctx.Context.Param("id")
	var article articles.Article
	model.Where("id = ?", id).Find(&article)
	ctx.ReturnSuccess(article)
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

// 文章详情
func (a *ArticleApi) GuestShow(ctx *route.Context) error {
	model := articles.ArticleDao.GetGuestModel(ctx.AppID)
	id := ctx.Context.Param("id")
	var article articles.Article
	model.Where("id = ?", id).Find(&article)
	ctx.ReturnSuccess(article)
	return nil
}

// 文章详情
func (a *ArticleApi) GuestList(ctx *route.Context) error {
	model := articles.ArticleDao.GetGuestModel(ctx.AppID)
	articlesList := []articles.Article{}
	model.Order("weight desc").Find(&articlesList)

	ctx.ReturnSuccess(articlesList)
	return nil
}
