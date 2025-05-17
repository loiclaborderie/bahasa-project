package models

import "gorm.io/gorm"

type Module struct {
	gorm.Model
	Title       string           `json:"title" gorm:"not null"`
	Description string           `json:"description"`
	Dialogues   []Dialogue       `json:"dialogues,omitempty" gorm:"foreignKey:ModuleID"`
	Vocabulary  []VocabularyItem `json:"vocabulary,omitempty" gorm:"many2many:module_vocabulary;"`
	// Grammar     []GrammarRule    `json:"grammar" gorm:"foreignKey:ModuleID"`
	// Tasks       []Task           `json:"tasks" gorm:"foreignKey:ModuleID"`
}
