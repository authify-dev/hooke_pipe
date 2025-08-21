package controllers

import (
	"common/domain/logger"
	"common/utils"
	"hook_pipe/internal/core/settings"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// HealthController estructura para manejar la ruta de Health
type HealthController struct {
}

// NewHealthController constructor para HealthController
func NewHealthController() *HealthController {
	return &HealthController{}
}

// GetHealth
func (c *HealthController) GetHealth(ctx *gin.Context) {

	entry := logger.FromContext(ctx.Request.Context())

	entry.Info("HealthController.GetHealth")

	customResponse := utils.Response[gin.H]{
		StatusCode: http.StatusOK,
		Data: gin.H{
			"status":      "ok",
			"message":     "El servicio está en línea y funcionando correctamente.",
			"timestamp":   time.Now().Unix(),
			"environment": settings.Settings.ENVIRONMENT,
			"app_name":    settings.Settings.APP_NAME,
		},
		Success: true,
	}

	// Se almacena el objeto para que el middleware lo procese
	ctx.JSON(http.StatusOK, customResponse)
}
