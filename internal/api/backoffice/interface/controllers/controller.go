package controllers

import "hook_pipe/internal/api/backoffice/app/services"

type HooksController struct {
	hooksService services.HooksService
}

func NewHooksController(hooksService services.HooksService) *HooksController {
	return &HooksController{hooksService: hooksService}
}
