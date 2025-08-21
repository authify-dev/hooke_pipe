package controllers

import (
	"hook_pipe/internal/api/hooks/app/services"
)

type HooksController struct {
	hooksService *services.HookPipeService
}

func NewHooksController(hooksService *services.HookPipeService) *HooksController {
	return &HooksController{hooksService: hooksService}
}
