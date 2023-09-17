package driver

import (
	"github.com/gin-gonic/gin"
	"github.com/takumi2786/zero-backend/internal/interface/controller"
)

// APIのルーティングを設定する
func SetRouting(
	gin *gin.Engine,
	lc *controller.LoginController,
	pc *controller.PostController,
) {
	group := gin.Group("")
	group.POST("/login", lc.Login)
	group.GET("/posts", pc.FindPosts)
	group.POST("/posts", pc.AddPost)
}
