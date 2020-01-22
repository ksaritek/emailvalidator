package register

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"strings"
)

type requireValidator struct {
	Validator *validator.Validate
}

func NewRequireValidator() *requireValidator {
	v := validator.New()
	return &requireValidator{
		Validator:v,
	}
}

func (d *requireValidator)Validate(p string) error {
	type emailRequest struct {
		Email string `json:"email" validate:"required"`
	}

	var e emailRequest
	if err := json.NewDecoder(strings.NewReader(p)).Decode(&e); err != nil {
		return err
	}

	return d.Validator.Struct(&e)
}