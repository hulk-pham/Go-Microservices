package api_test

import (
	"bytes"
	"encoding/json"
	"hulk/go-webservice/common"
	"hulk/go-webservice/infrastructure/persist"
	api "hulk/go-webservice/presentation/http"
	"hulk/go-webservice/presentation/http/auth"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginHandler(t *testing.T) {
	r := api.InitRouter()
	persist.InitDB()
	userMockCredential := auth.LoginRequest{
		Email:    "tanhunghue233@gmail.com",
		Password: "32323",
	}

	jsonValue, _ := json.Marshal(userMockCredential)
	req, _ := http.NewRequest("POST", "/api/auth/login", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var result common.JSONResult
	json.Unmarshal(w.Body.Bytes(), &result)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, result.Data)
}
