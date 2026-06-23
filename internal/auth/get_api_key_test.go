package auth

import (
	"testing"
	"net/http"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Add("Negroni", "salam/aleikum")
	headers.Add("Authorization", "ApiKey Oloilol")
 	got, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("error not nil: %v", err)
	}

	want := "Oloilol"
	if got != want {
		t.Fatalf("expected %v, got: %v", want, got)
	}
}
