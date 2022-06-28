package category

import (
	"errors"
	"github.com/glide-im/api/internal/dao/wrapper/category"
	"github.com/glide-im/api/internal/pkg/validate"
)

type CategoryStoreRequest struct {
	Name   string `json:"name" validate:"required,lte=15"`
	Weight int64  `json:"weight" validate:"required"`
	Icon   string `json:"icon"`
	ID     int
}

type CategoryUserRequest struct {
	CategoryIds []int64 `json:"category_ids" validate:"required"`
}

type Orders struct {
	ID     int `validate:"required"`
	Weight int `validate:"required"`
}

type CategoryOrderRequest struct {
	Orders []Orders `json:"orders"`
}

func (s *CategoryStoreRequest) Validate() error {
	if err := validate.ValidateHandle(s); err != nil {
		return validate.ValidateHandle(s)
	}

	var count int64
	name := s.Name
	model := category.CategoryDao.GetModel(1)
	model.Where("name = ?", name).Count(&count)

	if count > 0 {
		return errors.New("请不要重复创建分类:" + name)
	}

	return nil
}

func (s *CategoryUserRequest) Validate() error {
	return validate.ValidateHandle(s)
}
