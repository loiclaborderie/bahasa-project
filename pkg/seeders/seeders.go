package seeders

import (
	"fmt"
	"log"

	"github.com/loiclaborderie/bahasa-project/constants"
	"github.com/loiclaborderie/bahasa-project/internal/dialogue"
	"github.com/loiclaborderie/bahasa-project/internal/module"
	"github.com/loiclaborderie/bahasa-project/internal/user"
	"github.com/loiclaborderie/bahasa-project/internal/vocabulary"
	"github.com/loiclaborderie/bahasa-project/pkg/helper"
	"gorm.io/gorm"
)

// SeedUsers seeds default users
func SeedUsers(db *gorm.DB) error {
	credentials := []struct {
		Username string
		Email    string
		Password string
		Role     constants.Role
	}{
		{"admin", "admin@example.com", "admin123", constants.RoleAdmin},
		{"user", "user@example.com", "user123", constants.RoleUser},
	}

	for _, cred := range credentials {
		hashedPassword, err := helper.EncryptPassword(cred.Password)
		if err != nil {
			return fmt.Errorf("failed to hash password for %s: %w", cred.Email, err)
		}

		user := user.User{
			Username: cred.Username,
			Email:    cred.Email,
			Password: hashedPassword,
			Role:     cred.Role,
		}

		if err := db.Where("email = ?", user.Email).FirstOrCreate(&user).Error; err != nil {
			return fmt.Errorf("failed to seed user %s: %w", user.Email, err)
		}
	}

	return nil
}

// SeedModules seeds default modules
func SeedModules(db *gorm.DB) {
	modules := []module.Module{
		{Title: "Introduction", Description: "Getting started with Bahasa"},
		{Title: "Advanced", Description: "Advanced Bahasa lessons"},
	}
	for _, m := range modules {
		db.FirstOrCreate(&m, module.Module{Title: m.Title})
	}
}

// SeedDialogues seeds default dialogues
func SeedDialogues(db *gorm.DB) {
	dialogues := []dialogue.Dialogue{
		{Title: "Greetings", ModuleID: 1},
		{Title: "Shopping", ModuleID: 1},
	}
	for _, d := range dialogues {
		db.FirstOrCreate(&d, dialogue.Dialogue{Title: d.Title})
	}
}

// SeedDialogueLines seeds default dialogue lines
func SeedDialogueLines(db *gorm.DB) {
	lines := []dialogue.DialogueLine{
		{DialogueID: 1, Speaker: "A", Text: "Halo! Apa kabar ?", Translation: "Hello! How  are you ?"},
		{DialogueID: 1, Speaker: "B", Text: "Hai! Baik, terimah kasih.", Translation: "Hi! I'm good, thank you"},
	}
	for _, l := range lines {
		db.FirstOrCreate(&l, dialogue.DialogueLine{DialogueID: l.DialogueID, Text: l.Text})
	}
}

// SeedVocabularyItems seeds default vocabulary items
func SeedVocabularyItems(db *gorm.DB) {
	items := []vocabulary.VocabularyItem{
		{Term: "Halo", Definition: "Hello"},
		{Term: "Terima kasih", Definition: "Thank you"},
	}
	for _, i := range items {
		db.FirstOrCreate(&i, vocabulary.VocabularyItem{Term: i.Term})
	}
}

// SeedVocabularyLists seeds default vocabulary lists
func SeedVocabularyLists(db *gorm.DB) {
	// Find a user to be the creator (e.g., the first user)
	var creator user.User
	if err := db.First(&creator).Error; err != nil {
		log.Println("No user found to assign as creator for VocabularyLists. Please seed users first.")
		return
	}

	lists := []vocabulary.VocabularyList{
		{Title: "Basics", CreatedBy: creator.ID},
		{Title: "Greetings", CreatedBy: creator.ID},
	}
	for _, l := range lists {
		db.FirstOrCreate(&l, vocabulary.VocabularyList{Title: l.Title, CreatedBy: l.CreatedBy})
	}
}

// SeedAll runs all seeders
func SeedAll(db *gorm.DB) {
	log.Println("Seeding users...")
	SeedUsers(db)
	log.Println("Seeding modules...")
	SeedModules(db)
	log.Println("Seeding dialogues...")
	SeedDialogues(db)
	log.Println("Seeding dialogue lines...")
	SeedDialogueLines(db)
	log.Println("Seeding vocabulary items...")
	SeedVocabularyItems(db)
	log.Println("Seeding vocabulary lists...")
	SeedVocabularyLists(db)
	log.Println("Seeding completed.")
}
