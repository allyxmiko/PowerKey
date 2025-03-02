package validator

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtrans "github.com/go-playground/validator/v10/translations/zh"
	"strings"
)

var Validate *validator.Validate
var Trans ut.Translator

func Init() error {
	Validate = validator.New()
	uni := ut.New(zh.New())
	Trans, _ = uni.GetTranslator("zh")
	if err := zhtrans.RegisterDefaultTranslations(Validate, Trans); err != nil {
		return err
	}
	return nil
}

func Struct(s interface{}) (map[string]string, error) {
	errsMap := make(map[string]string)

	errs := Validate.Struct(s)
	if errs != nil {
		for _, e := range errs.(validator.ValidationErrors) {
			errsMap[strings.ToLower(e.Field())] = e.Translate(Trans)
		}
	}
	return errsMap, errs
}
