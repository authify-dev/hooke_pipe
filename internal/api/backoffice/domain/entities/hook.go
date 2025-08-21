package entities

import (
	"common/domain"
)

type HookEntity struct {
	domain.Entity
	Path        string `json:"path" binding:"required"`
	Port        int    `json:"port" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}
