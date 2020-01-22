package handler

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/ksaritek/emailvalidator/internal/domain"
	"github.com/ksaritek/emailvalidator/internal/register"
	"io/ioutil"
	"net/http"
)

func NewValidationHandler() http.Handler {
	d := register.NewDomainValidator()
	rexp := register.NewRegexpValidator()
	s := register.NewSmtpValidator()

	return validationHandler(d, rexp, s)
}

func validationHandler(validatorList ...register.Validator) http.HandlerFunc {
	required := register.NewRequireValidator()

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		p, _ := ioutil.ReadAll(r.Body)
		if err := required.Validate(r.Context(), string(p)); err != nil {
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		v := domain.Validation{Valid: true}
		v.Validators = &domain.Validators{}

		for _, validation := range validatorList {
			if err := validation.Validate(r.Context(), string(p)); err != nil {
				v.Valid = false

				for _, ve := range err.(validator.ValidationErrors) {
					switch ve.Tag() {
					case "regexp":
						v.Validators.Regexp = &domain.RegexpValidation{Status: false, Reason: "INVALID_EMAIL"}
					case "domain":
						v.Validators.Domain = &domain.DomainValidation{Status: false, Reason: "INVALID_TLD"}
					case "smtp":
						v.Validators.SMTP = &domain.SmtpValidation{Status: false, Reason: "UNABLE_TO_CONNECT"}
					}
				}
			}
		}

		if v.Validators.Regexp == nil {
			v.Validators.Regexp = &domain.RegexpValidation{Status: true}
		}

		if v.Validators.SMTP == nil {
			v.Validators.SMTP = &domain.SmtpValidation{Status: true}
		}

		if v.Validators.Domain == nil {
			v.Validators.Domain = &domain.DomainValidation{Status: true}
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(v); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
		}
	})
}
