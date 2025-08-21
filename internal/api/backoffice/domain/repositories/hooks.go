package repositories

import (
	"common/domain/criteria"
	"common/utils"
	"hook_pipe/internal/api/backoffice/domain/entities"
)

type HooksRepository interface {
	Save(hook entities.HookEntity) utils.Result[string]
	Matching(cr criteria.Criteria) ([]entities.HookEntity, error)
}
