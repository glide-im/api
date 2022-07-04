package app

import "github.com/glide-im/api/internal/pkg/validate"

type AppStoreRequest struct {
	Name string `json:"name" validate:"required,lte=100"`
	//License string `json:"license"`
	Logo  string `json:"logo" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Phone string `json:"phone" validate:"required"`
	AppId string `json:"app_id"`
	Host  string `json:"host"`
}

type Orders struct {
	ID     int `validate:"required"`
	Weight int `validate:"required"`
}

type AppOrderRequest struct {
	Orders []Orders `json:"orders"`
}

func (s *AppStoreRequest) Validate() error {
	if err := validate.ValidateHandle(s); err != nil {
		return err
	}
	return nil
}
