package register

import (
	"context"
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"net"
	"strings"
)

type domainValidator struct {
	Validator *validator.Validate
}

func NewDomainValidator() *domainValidator {
	v := validator.New()
	v.RegisterValidation("domain", func(fl validator.FieldLevel) bool {
		e := fl.Field().String()
		_, host := split(e)
		_, err := net.LookupMX(host)
		if err != nil {
			return false
		}

		return true
	})
	return &domainValidator{
		Validator: v,
	}
}

func (d *domainValidator) Validate(ctx context.Context, p string) error {
	type emailRequest struct {
		Email string `json:"email" validate:"domain"`
	}

	var e emailRequest
	if err := json.NewDecoder(strings.NewReader(p)).Decode(&e); err != nil {
		return err
	}

	return d.Validator.StructCtx(ctx, &e)
}
