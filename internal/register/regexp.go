package register

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
)

type regexpValidator struct {
	Validator *validator.Validate
}

func NewRegexpValidator() *regexpValidator {
	v := validator.New()
	var emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

	v.RegisterValidation("regexp", func(fl validator.FieldLevel) bool {
		if !emailRegexp.MatchString(fl.Field().String()) {
			return false
		}
		return true
	})
	return &regexpValidator{
		Validator:v,
	}
}

func (d *regexpValidator)Validate(p string) error {
	type emailRequest struct {
		Email string `json:"email" validate:"regexp"`
	}

	var e emailRequest
	if err := json.NewDecoder(strings.NewReader(p)).Decode(&e); err != nil {
		return err
	}

	return d.Validator.Struct(&e)
}