package testapiauth

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/joho/godotenv"
)

func performRequest(r http.Handler, method string, path string, body []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewReader(body))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func init() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func TestRegistration(t *testing.T) {
	// router := engine.GinEngine("test")
	// w := performRequest(router, "POST", "/api/v1/auth/reg", []byte{})
	// assert.Equal(t, http.StatusBadRequest, w.Code)

	// data := gin.H{
	// 	"login":    "login",
	// 	"password": "12345",
	// }
	// body, _ := json.Marshal(data)
	// w = performRequest(router, "POST", "/api/v1/auth/reg", body)
	// var m map[string]interface{}
	// err := json.NewDecoder(w.Body).Decode(&m)
	// user, ok := jwtService.ValidateToken(m["token"].(string))
	// assert.Equal(t, http.StatusOK, w.Code)
	// assert.Nil(t, err)
	// assert.Equal(t, ok, true)
	// assert.Equal(t, data["login"], user["login"])
}
