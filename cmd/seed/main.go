package main

import (
	"fmt"

	"github.com/loiclaborderie/bahasa-project/pkg/db"
	"github.com/loiclaborderie/bahasa-project/pkg/seeders"
)

func main() {
	database := db.Init()
	seeders.SeedAll(database)
	fmt.Println("Seeding completed.")
}
