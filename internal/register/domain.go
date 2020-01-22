package register

import (
	"github.com/go-playground/validator/v10"
	"net"
)

func DomainValidator(v *validator.Validate) {
	v.RegisterValidation("domain", func(fl validator.FieldLevel) bool {
		e := fl.Field().String()
		_, host := split(e)
		_, err := net.LookupMX(host)
		if err != nil {
			return false
		}

		return true
	})
}
