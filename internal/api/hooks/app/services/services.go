package services

import hooks_postgres "hook_pipe/internal/db/postgres/hooks"

type HookPipeService struct {
	hooksRepository *hooks_postgres.HooksPostgresRepository
}

func NewHookPipeService(hooksRepository *hooks_postgres.HooksPostgresRepository) *HookPipeService {
	return &HookPipeService{hooksRepository: hooksRepository}
}
