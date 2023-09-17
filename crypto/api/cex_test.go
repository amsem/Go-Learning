package api_test

import ("testing"
	"github.com/amsem/crypto/api"
	)

func TestAPICall(t *testing.T) {
	_, err := api.GetRate("")
	if err == nil {
		t.Error("Error was not found")
	}
}
