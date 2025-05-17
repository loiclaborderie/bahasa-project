package models

import "gorm.io/gorm"

// Dialogue represents a conversation used for language learning
type Dialogue struct {
	gorm.Model
	ModuleID    uint   `json:"module_id" gorm:"not null"`
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description"`
	// Relations
	Module Module         `json:"-" gorm:"foreignKey:ModuleID"`
	Lines  []DialogueLine `json:"lines,omitempty" gorm:"foreignKey:DialogueID"`
}

// DialogueLine represents a single line in a dialogue
type DialogueLine struct {
	gorm.Model
	DialogueID  uint   `json:"dialogue_id" gorm:"not null"`
	Speaker     string `json:"speaker" gorm:"not null"`
	Text        string `json:"text" gorm:"not null"`
	Translation string `json:"translation"`
	AudioURL    string `json:"audio_url"`
	OrderIndex  int    `json:"order_index" gorm:"not null"`
	// Relations
	Dialogue Dialogue `json:"-" gorm:"foreignKey:DialogueID"`
}
