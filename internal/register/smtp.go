package register

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net"
	"strings"
	"time"
)


type smtpValidator struct {
	Validator *validator.Validate
}

func NewSmtpValidator() *smtpValidator {
	v := validator.New()
	const forceDisconnectAfter = time.Second * 5

	v.RegisterValidation("smtp", func(fl validator.FieldLevel) bool {
		e := fl.Field().String()
		_, host := split(e)
		mx, err := net.LookupMX(host)
		if err != nil {
			return false
		}

		addr := fmt.Sprintf("%s:%d", strings.TrimRight(mx[0].Host, "."), 25)
		if host == "gmail.com" {
			addr = "smtp.gmail.com:587"
		}

		client, err := dialTimeout(addr, forceDisconnectAfter)
		if err != nil {
			return false
		}
		client.Quit()

		return true
	})

	return &smtpValidator{
		Validator:v,
	}
}

func (d *smtpValidator)Validate(p string) error {
	type emailRequest struct {
		Email string `json:"email" validate:"smtp"`
	}

	var e emailRequest
	if err := json.NewDecoder(strings.NewReader(p)).Decode(&e); err != nil {
		return err
	}

	return d.Validator.Struct(&e)
}