package hooks

import (
	"hook_pipe/internal/api/hooks/app/services"
	"hook_pipe/internal/api/hooks/interface/controllers"
	hooks_postgres "hook_pipe/internal/db/postgres/hooks"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

func SetupHookPipeModule(r *gin.Engine, db *gorm.DB) {

	// Repositories
	hooksRepository := hooks_postgres.NewHooksPostgresRepository(db)

	// Services
	hooksService := services.NewHookPipeService(hooksRepository)

	// Controllers
	hooksController := controllers.NewHooksController(hooksService)

	// Routes
	hooksGroup := r.Group("/webhook")
	hooksGroup.Any("/:vendor/*path", hooksController.Pipe)
}
