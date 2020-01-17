/*
@Time : 2020/1/16 23:35
@Author : Minus4
*/
package validate

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	enTranslations "gopkg.in/go-playground/validator.v9/translations/en"
	"m4-im/pkg/util"
)

var (
	Uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func SetupValidator(v *validator.Validate) {
	validate = v
	engLocal := en.New()
	Uni = ut.New(engLocal, engLocal)
	trans, _ := Uni.GetTranslator("en")

	// Register en trans
	_ = enTranslations.RegisterDefaultTranslations(validate, trans)

	// Register email tag
	_ = validate.RegisterValidation("email", emailValidation)
}

func emailValidation(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	return util.VerifyEmailFormat(email)
}
