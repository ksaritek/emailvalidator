package handler

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	"github.com/ksaritek/emailvalidator/internal/domain"
	"github.com/ksaritek/emailvalidator/internal/register"
	"net/http"
)

type emailRequest struct {
	Email string `json:"email" validate:"required,regexp,domain,smtp"`
}

func NewValidationHandler() http.Handler {
	validate := validator.New()

	register.RegexpValidator(validate)
	register.DomainValidator(validate)
	register.SMTPValidator(validate)

	return validationHandler(validate)
}

func validationHandler(validate *validator.Validate) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var e emailRequest
		if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		v := domain.Validation{Valid: true}
		v.Validators = &domain.Validators{
			Regexp: &domain.RegexpValidation{Status: true},
			Domain: &domain.DomainValidation{Status: true},
			SMTP:   &domain.SmtpValidation{Status: true},
		}

		err := validate.Struct(e)
		if err != nil {
			v.Valid = false

			for _, ve := range err.(validator.ValidationErrors) {
				switch ve.Tag() {
				case "regexp":
					v.Validators.Regexp = &domain.RegexpValidation{Status: false, Reason: "INVALID_EMAIL"}
				case "domain":
					v.Validators.Domain = &domain.DomainValidation{Status: false, Reason: "INVALID_TLD"}
				case "smtp":
					v.Validators.SMTP = &domain.SmtpValidation{Status: false, Reason: "UNABLE_TO_CONNECT"}
				case "required":
					v.Validators = nil
				}
			}
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(v); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
		}
	})
}
