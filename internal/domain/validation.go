package domain

type Validation struct {
	Valid      bool        `json:"valid"`
	Validators *Validators `json:"validators,omitempty"`
}

type Validators struct {
	Regexp *RegexpValidation `json:"regexp,omitempty"`
	Domain *DomainValidation `json:"domain,omitempty"`
	SMTP   *SmtpValidation   `json:"smtp,omitempty"`
}

type RegexpValidation struct {
	Status bool   `json:"valid"`
	Reason string `json:"reason,omitempty"`
}

type DomainValidation struct {
	Status bool   `json:"valid"`
	Reason string `json:"reason,omitempty"`
}

type SmtpValidation struct {
	Status bool   `json:"valid"`
	Reason string `json:"reason,omitempty"`
}
