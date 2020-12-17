package test

import (
	"github.com/stretchr/testify/assert"
	"go-index/packages/gin_series/01_hello_gin/initRouter"
	"net/http"
	"net/http/httptest"
	"testing"
)

/**
单元测试
 */
func TestIndexGetRouter(t *testing.T) {
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "hello gin", w.Body.String())
}
