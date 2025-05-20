package vocabulary

import (
	"github.com/loiclaborderie/bahasa-project/internal/user"
	"gorm.io/gorm"
)

type VocabularyItem struct {
	gorm.Model
	Term         string `json:"term" gorm:"not null"`
	Definition   string `json:"definition" gorm:"not null"`
	Example      string `json:"example"`
	PartOfSpeech string `json:"part_of_speech"` // Noun, Verb, etc.
	Difficulty   string `json:"difficulty"`     // Easy, Medium, Hard
	AudioURL     string `json:"audio_url"`
	ImageURL     string `json:"image_url"`
}

type VocabularyList struct {
	gorm.Model
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description"`
	CreatedBy   uint   `json:"created_by" gorm:"not null"`
	// Relations
	Creator user.User        `json:"-" gorm:"foreignKey:CreatedBy"`
	Items   []VocabularyItem `json:"items,omitempty" gorm:"many2many:list_vocabulary;"`
}
