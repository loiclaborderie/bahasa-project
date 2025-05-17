package seeders

import (
	"log"

	"github.com/loiclaborderie/bahasa-project/models"
	"gorm.io/gorm"
)

// SeedUsers seeds default users
func SeedUsers(db *gorm.DB) {
	users := []models.User{
		{Username: "admin", Email: "admin@example.com", Password: "admin123"},
		{Username: "user", Email: "user@example.com", Password: "user123"},
	}
	for _, u := range users {
		db.FirstOrCreate(&u, models.User{Email: u.Email})
	}
}

// SeedModules seeds default modules
func SeedModules(db *gorm.DB) {
	modules := []models.Module{
		{Title: "Introduction", Description: "Getting started with Bahasa"},
		{Title: "Advanced", Description: "Advanced Bahasa lessons"},
	}
	for _, m := range modules {
		db.FirstOrCreate(&m, models.Module{Title: m.Title})
	}
}

// SeedDialogues seeds default dialogues
func SeedDialogues(db *gorm.DB) {
	dialogues := []models.Dialogue{
		{Title: "Greetings", ModuleID: 1},
		{Title: "Shopping", ModuleID: 1},
	}
	for _, d := range dialogues {
		db.FirstOrCreate(&d, models.Dialogue{Title: d.Title})
	}
}

// SeedDialogueLines seeds default dialogue lines
func SeedDialogueLines(db *gorm.DB) {
	lines := []models.DialogueLine{
		{DialogueID: 1, Speaker: "A", Text: "Halo! Apa kabar ?", Translation: "Hello! How  are you ?"},
		{DialogueID: 1, Speaker: "B", Text: "Hai! Baik, terimah kasih.", Translation: "Hi! I'm good, thank you"},
	}
	for _, l := range lines {
		db.FirstOrCreate(&l, models.DialogueLine{DialogueID: l.DialogueID, Text: l.Text})
	}
}

// SeedVocabularyItems seeds default vocabulary items
func SeedVocabularyItems(db *gorm.DB) {
	items := []models.VocabularyItem{
		{Term: "Halo", Definition: "Hello"},
		{Term: "Terima kasih", Definition: "Thank you"},
	}
	for _, i := range items {
		db.FirstOrCreate(&i, models.VocabularyItem{Term: i.Term})
	}
}

// SeedVocabularyLists seeds default vocabulary lists
func SeedVocabularyLists(db *gorm.DB) {
	// Find a user to be the creator (e.g., the first user)
	var creator models.User
	if err := db.First(&creator).Error; err != nil {
		log.Println("No user found to assign as creator for VocabularyLists. Please seed users first.")
		return
	}

	lists := []models.VocabularyList{
		{Title: "Basics", CreatedBy: creator.ID},
		{Title: "Greetings", CreatedBy: creator.ID},
	}
	for _, l := range lists {
		db.FirstOrCreate(&l, models.VocabularyList{Title: l.Title, CreatedBy: l.CreatedBy})
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
