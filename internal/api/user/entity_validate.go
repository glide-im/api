package user

import (
	"errors"
	"github.com/glide-im/api/internal/pkg/validate"
)

func (request *UpdateProfileRequest) Validate() error {
	if err := validate.ValidateHandle(request); err != nil {
		return err
	}

	if len(request.Password) > 0 {
		errs := validate.Validate.Var(request.Password, "gte=6,lte=16")
		if errs != nil {
			return errors.New("密码格式不符合")
		}
	}

	return nil
}
