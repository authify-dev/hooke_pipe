package hooks_postgres

import (
	postgres "common/infrastructure/db/gorm"
	"hook_pipe/internal/api/backoffice/domain/entities"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// --------------------------------
// INFRASTRUCTURE
// --------------------------------
// User Model
// --------------------------------

// CodeModel utiliza Model parametrizado con User.
type HookModel struct {
	postgres.Model[entities.HookEntity]
	Path        string `gorm:"type:varchar(255);not null;" json:"path"`
	Port        int    `gorm:"type:int;not null" json:"port"`
	Name        string `gorm:"type:varchar(255);not null;" json:"name"`
	Description string `gorm:"type:varchar(255);not null;" json:"description"`
}

func (HookModel) TableName() string {
	return "hooks"
}

func (c HookModel) GetID() string {
	return c.ID.String()
}

func (m *HookModel) BeforeCreate(tx *gorm.DB) (err error) {
	m.ID = uuid.New()
	m.CreatedAt = time.Now().UTC()
	m.UpdatedAt = time.Now().UTC()
	m.IsRemoved = false
	return m.Model.BeforeCreate(tx)
}
