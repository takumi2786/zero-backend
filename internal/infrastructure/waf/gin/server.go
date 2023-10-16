package gin

import (
	"fmt"
	"net/http"
	"time"

	ginzap "github.com/gin-contrib/zap"
	gin "github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/takumi2786/zero-backend/internal/interface/controller"
	"github.com/takumi2786/zero-backend/internal/util"
	"go.uber.org/zap"
)

type GinApp struct {
	cfg    *util.Config
	logger *zap.Logger
	db     *sqlx.DB
}

func NewGinApp(cfg *util.Config, logger *zap.Logger, db *sqlx.DB) *GinApp {
	return &GinApp{cfg: cfg, logger: logger, db: db}
}

func (ga *GinApp) Run(lc *controller.LoginController) error {
	if ga.cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	app := gin.Default()
	app.Use(ginzap.Ginzap(ga.logger, time.RFC3339, true))

	/*
		Routing
	*/
	// setup router
	group := app.Group("")
	group.POST("/login", ga.getLoginHandler(lc))

	ga.logger.Info("Start Server")
	err := app.Run(fmt.Sprintf(":%d", ga.cfg.Port))
	return err
}

func (ga *GinApp) getLoginHandler(lc *controller.LoginController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 入力値を取得
		var loginRequest controller.LoginRequestBody
		if err := ctx.ShouldBindJSON(&loginRequest); err != nil {
			ctx.JSON(http.StatusBadRequest, http.StatusText(http.StatusBadRequest))
			lc.Logger.Error("failed to parse request", zap.Error(err))
			return
		}
		successResp, errResp := lc.Login(loginRequest)
		if errResp != nil {
			ctx.JSON(errResp.Code, errResp.Message)
			return
		}
		fmt.Print(errResp)
		ctx.JSON(successResp.Code, successResp.Body)
		return
	}
}
