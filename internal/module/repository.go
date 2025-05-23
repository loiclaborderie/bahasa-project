package module

import (
	"context"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type ModuleRepository struct {
	db *gorm.DB
}

func NewModuleRepository(db *gorm.DB) *ModuleRepository {
	return &ModuleRepository{db: db}
}

func (r *ModuleRepository) GetAll(ctx context.Context) (*[]Module, error) {
	var module []Module
	result := r.db.WithContext(ctx).Find(&module)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("no modules found: %w", result.Error)
		}
		return nil, fmt.Errorf("database error: %w", result.Error)
	}
	return &module, nil
}
func (r *ModuleRepository) GetVisible(ctx context.Context) (*[]Module, error) {
	var module []Module
	result := r.db.WithContext(ctx).Where("visible = ?", true).Find(&module)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("no modules found: %w", result.Error)
		}
		return nil, fmt.Errorf("database error: %w", result.Error)
	}
	return &module, nil
}

func (r *ModuleRepository) GetByID(ctx context.Context, id int) (*Module, error) {
	var module Module
	err := r.db.WithContext(ctx).
		Model(&Module{}).
		Where("modules.id = ?", id).
		Preload("Vocabulary").
		Preload("Dialogues", func(db *gorm.DB) *gorm.DB {
			return db.
				Joins("JOIN dialogue_lines ON dialogue_lines.dialogue_id = dialogues.id").
				Group("dialogues.id") // prevent duplicate rows
		}).
		First(&module).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("module not found: %w", err)
		}
		return nil, fmt.Errorf("database error: %w", err)
	}
	return &module, nil
}

func (r *ModuleRepository) Create(ctx context.Context, module *Module) (*Module, error) {
	result := r.db.WithContext(ctx).Create(&module)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("module not found: %w", result.Error)
		}
		return nil, fmt.Errorf("database error: %w", result.Error)
	}
	return module, nil
}
