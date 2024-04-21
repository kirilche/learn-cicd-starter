package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	header := http.Header(map[string][]string{
		"Authorization": {""},
	})
	res, err := GetAPIKey(header)
	if err == nil {
		t.Error("there should be an error", res, err)
	}
}
