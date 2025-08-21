package services

import "hook_pipe/internal/api/backoffice/domain/repositories"

type HooksService struct {
	hooksRepository repositories.HooksRepository
}

func NewHooksService(hooksRepository repositories.HooksRepository) *HooksService {
	return &HooksService{hooksRepository: hooksRepository}
}
