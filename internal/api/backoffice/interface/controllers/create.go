package controllers

import (
	"common/domain/customctx"
	"common/domain/logger"
	"hook_pipe/internal/api/backoffice/interface/dtos"

	"common/interface/ppdtos"

	"github.com/gin-gonic/gin"
)

func (c *HooksController) Create(ctx *gin.Context) {

	entry := logger.FromContext(ctx.Request.Context())

	entry.Info("Creating hook pipe")

	cc := customctx.NewCustomContext(ctx.Request.Context())

	dto := ppdtos.GetDTO[dtos.CreateHookDTO](ctx, cc)

	response := c.hooksService.Create(cc, dto.ToCommand())

	ctx.JSON(response.StatusCode, response.ToMapWithCustomContext(cc))
}
