package user

import (
	"errors"
	"github.com/glide-im/api/internal/pkg/validate"
	"github.com/go-playground/validator/v10"
)

func validateHandle(request interface{}) error {
	err := validate.Validate.Struct(request)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		_errors := errs.Translate(*validate.Translator)
		for _, _error := range _errors {
			return errors.New(_error)
		}
	}
	return nil
}

func (request *UpdateProfileRequest) Validate() error {
	if err := validateHandle(request); err != nil {
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
