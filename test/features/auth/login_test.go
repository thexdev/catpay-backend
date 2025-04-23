package auth_test

import (
	"catpay/internal/infra/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	t.Run("GET /", func(t *testing.T) {
		app := http.New().Bootstrap()

		req := httptest.NewRequest("GET", "/", nil)

		res, err := app.Test(req)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, 404, res.StatusCode)
	})
}
