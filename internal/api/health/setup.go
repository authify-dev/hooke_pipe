package health

import (
	"hook_pipe/internal/api/health/interface/controllers"
	"hook_pipe/internal/core/settings"

	"github.com/gin-gonic/gin"
)

func SetupHealthModule(app *gin.Engine) {

	healthController := controllers.NewHealthController()

	// Rutas de health
	health := app.Group(settings.Settings.ROOT_PATH + "/health")

	health.GET("", healthController.GetHealth)
}
