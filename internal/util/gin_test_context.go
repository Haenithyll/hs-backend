package util

import (
	"bytes"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func NewTestGinContext(method, path string, body string, userId string) (*gin.Context, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	req := httptest.NewRequest(method, path, bytes.NewBufferString(body))
	req.Header.Set("Content-Type", "application/json")

	gin.SetMode(gin.ReleaseMode)
	c, _ := gin.CreateTestContext(w)

	c.Request = req

	c.Set("user_id", userId)

	return c, w
}
