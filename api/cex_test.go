package api_test

import (
	"testing"

	"github.com/developertomek/crypto/api"
)

func TestAPICall(t *testing.T) {
	_, err := api.GetRate("")
	if err == nil {
		t.Errorf("Error was not found")
	}
}
