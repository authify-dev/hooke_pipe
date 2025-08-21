package server

import (
	"fmt"
	"hook_pipe/internal/api/health"
	"hook_pipe/internal/api/hooks"
	"hook_pipe/internal/core/settings"
	"hook_pipe/internal/router"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-gonic/gin"
)

func Run() {

	app := setUpRouter()

	if _, inLambda := os.LookupEnv("AWS_LAMBDA_FUNCTION_NAME"); inLambda {
		fmt.Println("Running in Lambda")
		lambda.Start(ginadapter.NewV2(app).ProxyWithContext)
		return
	}

	app.Run(fmt.Sprintf(":%d", settings.Settings.PORT))
}

func setUpRouter() *gin.Engine {
	app := router.NewRouter()

	// db, err := gorm.Open(sqlite.Open("hooks.db"), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// }

	health.SetupHealthModule(app)
	//backoffice.SetupBackofficeModule(app, db)
	hooks.SetupHookPipeModule(app)

	return app
}
