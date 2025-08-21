package services

import (
	"common/domain/customctx"
	"common/domain/logger"
	"common/utils"
	"hook_pipe/internal/api/backoffice/domain/commands"
	"hook_pipe/internal/api/backoffice/domain/entities"
	"net/http"

	"github.com/google/uuid"
)

func (s *HooksService) Create(cc *customctx.CustomContext, hook commands.CreateHookCommand) utils.Response[entities.HookEntity] {

	entry := logger.FromContext(cc.Context())

	entry.Info("Creating hook pipe")

	hookEntity := hook.ToEntity()

	entry.Info("Hook pipe created", hookEntity)

	hookEntity.ID = uuid.New().String()

	result := s.hooksRepository.Save(hookEntity)

	if result.Err != nil {
		entry.Error("Error creating hook pipe", result.Err)
		return utils.Response[entities.HookEntity]{
			StatusCode: http.StatusInternalServerError,
			Error:      cc.NewError(result.Err),
			Success:    false,
		}
	}

	hookEntity.ID = result.Data

	return utils.Response[entities.HookEntity]{Data: hookEntity, StatusCode: http.StatusCreated, Success: true}
}
