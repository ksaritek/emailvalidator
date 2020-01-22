package main

import (
	"fmt"
	"github.com/ksaritek/emailvalidator/internal/handler"
	"log"
	"net/http"
	"os"
)

func main() {
	p := fmt.Sprintf(":%s", os.Getenv("PORT"))
	if p == ":" {
		p = ":8080"
		log.Printf("fallback to default port %s\n", p)
	}

	http.Handle("/email/validate", handler.NewValidationHandler())
	log.Printf("Listening on %v\n", p)
	log.Fatal(http.ListenAndServe(p, nil))
}
