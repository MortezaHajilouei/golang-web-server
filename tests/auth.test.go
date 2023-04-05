package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignup(t *testing.T) {
	req, _ := http.NewRequest("POST", "/user/register/", nil)
	w := httptest.NewRecorder()
	// r.ServeHTTP(w, req)

	// var companies []User
	// json.Unmarshal(w.Body.Bytes(), &companies)

	// assert.Equal(t, http.StatusOK, w.Code)
	// assert.NotEmpty(t, companies)
}
