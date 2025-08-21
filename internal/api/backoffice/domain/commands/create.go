package commands

import "hook_pipe/internal/api/backoffice/domain/entities"

type CreateHookCommand struct {
	Path        string `json:"path" binding:"required"`
	Port        int    `json:"port" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

func (c *CreateHookCommand) ToEntity() entities.HookEntity {
	return entities.HookEntity{
		Path:        c.Path,
		Port:        c.Port,
		Name:        c.Name,
		Description: c.Description,
	}
}
