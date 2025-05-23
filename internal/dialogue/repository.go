package dialogue

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type DialogueRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *DialogueRepository {
	return &DialogueRepository{db: db}
}

func (d *DialogueRepository) GetAvailableSpeakers(c context.Context, dialogueId int) ([]string, error) {
	var speakers []string
	err := d.db.WithContext(c).
		Model(&DialogueLine{}).
		Select("DISTINCT speaker").
		Where("dialogue_id = ?", dialogueId).
		Order("speaker").
		Pluck("speaker", &speakers).Error
	return speakers, err
}

func (d *DialogueRepository) GetDialogue(c context.Context, dialogueId int) (*Dialogue, error) {
	var dialogue Dialogue
	err := d.db.WithContext(c).
		Preload("Lines").
		First(&dialogue, dialogueId).
		Error

	if err != nil {
		return nil, err
	}

	if len(dialogue.Lines) == 0 {
		return nil, fmt.Errorf("dialogue %d has no associated lines", dialogueId)
	}

	return &dialogue, nil
}
