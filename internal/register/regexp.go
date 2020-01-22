package register

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

func RegexpValidator(v *validator.Validate) {
	var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	v.RegisterValidation("regexp", func(fl validator.FieldLevel) bool {
		if !emailRegexp.MatchString(fl.Field().String()) {
			return false
		}
		return true
	})
}
