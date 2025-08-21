package dtos

import "hook_pipe/internal/api/backoffice/domain/commands"

type CreateHookDTO struct {
	Path        string `json:"path" binding:"required"`
	Port        int    `json:"port" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

func (dto CreateHookDTO) Validate() error {
	return nil
}

func (dto CreateHookDTO) ToCommand() commands.CreateHookCommand {
	return commands.CreateHookCommand{
		Path:        dto.Path,
		Port:        dto.Port,
		Name:        dto.Name,
		Description: dto.Description,
	}
}
