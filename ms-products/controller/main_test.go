package controller

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

// type TestEnv struct {
// 	ctrl     *gomock.Controller
// 	store    *mockdb.MockStore
// 	server   *Server
// 	recorder *httptest.ResponseRecorder
// 	url      string
// }

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}
