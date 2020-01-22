package register

import (
	"context"
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
		Validator: v,
	}
}

func (d *requireValidator) Validate(ctx context.Context, p string) error {
	type emailRequest struct {
		Email string `json:"email" validate:"required"`
	}

	var e emailRequest
	if err := json.NewDecoder(strings.NewReader(p)).Decode(&e); err != nil {
		return err
	}

	return d.Validator.StructCtx(ctx, &e)
}
