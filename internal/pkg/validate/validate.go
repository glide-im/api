package validate

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	uni        *ut.UniversalTranslator
	Validate   *validator.Validate
	Translator *ut.Translator
)

func Init() {
	zh := zh.New()
	uni = ut.New(zh, zh)
	trans, _ := uni.GetTranslator("zh")
	Translator = &trans

	Validate = validator.New()
	_ = zh_translations.RegisterDefaultTranslations(Validate, *Translator)
}
