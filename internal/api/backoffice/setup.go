package backoffice

import (
	"hook_pipe/internal/api/backoffice/app/services"
	"hook_pipe/internal/api/backoffice/interface/controllers"
	hooks_postgres "hook_pipe/internal/db/postgres/hooks"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupBackofficeModule(router *gin.Engine, db *gorm.DB) {

	// Repositories
	hooksRepository := hooks_postgres.NewHooksPostgresRepository(db)

	// Services
	hooksService := services.NewHooksService(hooksRepository)

	//Controllers
	hooksController := controllers.NewHooksController(*hooksService)

	// Routes
	router.POST("/hooks", hooksController.Create)
}
