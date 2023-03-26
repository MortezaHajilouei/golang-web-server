package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/MortezaHajilouei/golang-web-server/controllers"
	"github.com/stretchr/testify/assert"
	"github.com/wpcodevo/golang-gorm-postgres/controllers"
)

func TestGetCompaniesHandler(t *testing.T) {
	r := SetUpRouter()
	r.GET("/users/all", controllers.UserController)
	req, _ := http.NewRequest("GET", "/companies", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var companies []User
	json.Unmarshal(w.Body.Bytes(), &companies)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, companies)
}
