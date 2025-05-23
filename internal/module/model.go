package module

import (
	"github.com/loiclaborderie/bahasa-project/internal/dialogue"
	"github.com/loiclaborderie/bahasa-project/internal/vocabulary"
	"gorm.io/gorm"
)

type Module struct {
	gorm.Model
	Title       string                      `json:"title" gorm:"not null"`
	Description string                      `json:"description"`
	Visible     bool                        `json:"visible" gorm:"default:false"`
	CoverURL    string                      `json:"cover_url"`
	Dialogues   []dialogue.Dialogue         `json:"dialogues,omitempty" gorm:"foreignKey:ModuleID"`
	Vocabulary  []vocabulary.VocabularyItem `json:"vocabulary,omitempty" gorm:"many2many:module_vocabulary;"`
	// Grammar     []GrammarRule    `json:"grammar" gorm:"foreignKey:ModuleID"`
	// Tasks       []Task           `json:"tasks" gorm:"foreignKey:ModuleID"`
}
