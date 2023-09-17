package driver

import (
	"github.com/gin-gonic/gin"
	"github.com/takumi2786/zero-backend/internal/interface/controller"
)

// APIのルーティングを設定する
func SetRouting(
	gin *gin.Engine,
	lc *controller.LoginControler,
) {
	group := gin.Group("")
	group.POST("/login", lc.Login)
}
