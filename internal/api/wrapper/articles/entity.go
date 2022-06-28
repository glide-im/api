package articles

import "github.com/glide-im/api/internal/pkg/validate"

type ArticlesStoreRequest struct {
	Title     string `json:"title" validate:"required,lte=100"`
	PublishAT string `json:"publish_at"`
	Content   string `json:"content" validate:"required"`
	Weight    int64  `json:"weight" validate:"required"`
	ID        int
}

type Orders struct {
	ID     int `validate:"required"`
	Weight int `validate:"required"`
}

type ArticlesOrderRequest struct {
	Orders []Orders `json:"orders"`
}

func (s *ArticlesStoreRequest) Validate() error {
	return validate.ValidateHandle(s)
}
