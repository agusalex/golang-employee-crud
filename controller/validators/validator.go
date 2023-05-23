package validators

import "github.com/go-playground/validator/v10"

var Validate = validator.New()

func init() {
	Validate.RegisterValidation("gt", gt)
}

var gt validator.Func = func(fl validator.FieldLevel) bool {
	value, _ := fl.Field().Interface().(int)
	return value >= 0
}
