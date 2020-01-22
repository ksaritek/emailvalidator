package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/ksaritek/emailvalidator/internal/domain"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func Test_EmailValidation(t *testing.T) {
	tests := []struct {
		name       string
		email      string
		want       string
		statusCode int
	}{
		{
			name:       `valid email`,
			email:      `{"email":"xxx@gmail.com"}`,
			want:       `{"valid":true,"validators":{"regexp":{"valid":true},"domain":{"valid":true},"smtp":{"valid":true}}}`,
			statusCode: http.StatusOK,
		},
		{
			name:       `invalid domain & smtp connect failure`,
			email:      `{"email":"xxx@zhuu.zu"}`,
			want:       `{"valid":false,"validators":{"regexp":{"valid":true},"domain":{"valid":false,"reason":"INVALID_TLD"},"smtp":{"valid":false,"reason":"UNABLE_TO_CONNECT"}}}`,
			statusCode: http.StatusOK,
		},
		{
			name:       `invalid regexp`,
			email:      `{"email":"£££testinvalid@gmail.com"}`,
			want:       `{"valid":false,"validators":{"regexp":{"valid":false,"reason":"INVALID_EMAIL"},"domain":{"valid":true},"smtp":{"valid":true}}}`,
			statusCode: http.StatusOK,
		},
		{
			name:       `invalid regexp & domain & smtp connect failure`,
			email:      `{"email":"£££testinvalid@zhuuu.zu"}`,
			want:       `{"valid":false,"validators":{"regexp":{"valid":false,"reason":"INVALID_EMAIL"},"domain":{"valid":false,"reason":"INVALID_TLD"},"smtp":{"valid":false,"reason":"UNABLE_TO_CONNECT"}}}`,
			statusCode: http.StatusOK,
		},
		{
			name:       `invalid payload`,
			email:      `{"invalidp":"£££testinvalid@zhuuu.zu"}`,
			want:       "",
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := httptest.NewServer(NewValidationHandler())
			defer ts.Close()

			res, err := http.Post(fmt.Sprintf("%s/email/validate", ts.URL), "application/json", strings.NewReader(tt.email))
			if err != nil {
				log.Fatal(err)
			}

			if res.StatusCode == http.StatusBadRequest && res.StatusCode == tt.statusCode {
				return
			}

			var got domain.Validation
			if err != json.NewDecoder(res.Body).Decode(&got) {
				t.Fatal(err)
			}

			var want domain.Validation
			if err != json.NewDecoder(strings.NewReader(tt.want)).Decode(&want) {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(got, want) {
				var buf bytes.Buffer
				json.NewEncoder(&buf).Encode(got)
				t.Errorf("Email Validation Response - %v case => got %v, want %v", tt.name, buf.String(), tt.want)
			}
		})
	}
}
