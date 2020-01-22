package register

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"net"
	"strings"
	"time"
)


func SMTPValidator(v *validator.Validate) {
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
}
