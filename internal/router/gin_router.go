package router

import (
	"hook_pipe/internal/common/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	r.RedirectTrailingSlash = true
	r.RedirectFixedPath = true
	r.Use(gin.Recovery())

	r.Use(middlewares.RequestLogMiddleware())

	r.Use(cors.Default())

	return r
}
