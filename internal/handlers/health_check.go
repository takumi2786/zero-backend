package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckResponse struct {
	Message string `json:"message"`
}

type HealthCheck struct{}

func (hc *HealthCheck) ServeHTTP(ctx *gin.Context) {
	resp := HealthCheckResponse{
		Message: "ok",
	}
	ctx.JSON(http.StatusOK, resp)
}

func NewHealthCheckHandler() *HealthCheck {
	return &HealthCheck{}
}
