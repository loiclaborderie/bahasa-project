package module

import (
	"context"
)

type ModuleService struct {
	repo *ModuleRepository
}

func NewModuleService(repo *ModuleRepository) *ModuleService {
	return &ModuleService{repo: repo}
}

func (s *ModuleService) GetAll(ctx context.Context) (*[]Module, error) {
	module, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return module, nil
}

func (s *ModuleService) GetVisible(ctx context.Context) (*[]Module, error) {
	module, err := s.repo.GetVisible(ctx)
	if err != nil {
		return nil, err
	}

	return module, nil
}

func (s *ModuleService) GetByID(ctx context.Context, id int) (*Module, error) {
	module, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return module, nil
}

func (s *ModuleService) Create(ctx context.Context, module *Module) (*Module, error) {
	moduleCreated, err := s.repo.Create(ctx, module)
	if err != nil {
		return nil, err
	}

	return moduleCreated, nil
}
