package hooks_postgres

import (
	"common/domain/criteria"
	postgres "common/infrastructure/db/gorm"
	"hook_pipe/internal/api/backoffice/domain/entities"

	"gorm.io/gorm"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// Role Postgres Repository
// --------------------------------

type HooksPostgresRepository struct {
	postgres.PostgresRepository[entities.HookEntity, HookModel]
}

func NewHooksPostgresRepository(connection *gorm.DB) *HooksPostgresRepository {
	return &HooksPostgresRepository{
		PostgresRepository: postgres.PostgresRepository[entities.HookEntity, HookModel]{
			Connection: connection,
		},
	}
}

func (r *HooksPostgresRepository) Matching(cr criteria.Criteria) ([]entities.HookEntity, error) {

	model := &HookModel{}

	return r.MatchingLow(cr, model)
}
