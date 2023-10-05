package config

import (
	"fmt"
	"github.com/danielsantil/weather-api/internal/models/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

// AddDatabase connects to postgres database specified in connection string,
// applies migrations and seeds the database.
func AddDatabase(connString string) *gorm.DB {
	db, dbErr := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if dbErr != nil {
		log.Fatal("Can't connect to database: ", dbErr)
	}

	log.Println("Migrating database...")
	mErr := db.AutoMigrate(&database.Weather{}, &database.System{},
		&database.CityForecast{}, &database.Forecast{}, &database.City{},
		&database.Seeding{})
	if mErr != nil {
		log.Fatal("Can't migrate database: ", mErr)
	}
	log.Println("Migration complete")

	seedData(db)

	return db
}

// seedData Seeds database
//
// See sql/seeds directory for seeding files
func seedData(db *gorm.DB) {
	log.Println("Started seeding SQL files...")
	path := "internal/sql/seeds"
	entries, dirErr := os.ReadDir(path)
	if dirErr != nil {
		log.Println("Error reading seeding directory", dirErr)
		return
	}

	var seeding []database.Seeding
	db.Find(&seeding)
	var count int

	for _, file := range entries {
		if seedingExists(seeding, file.Name()) {
			continue
		}

		fileBytes, fErr := os.ReadFile(fmt.Sprintf("%s/%s", path, file.Name()))
		if fErr != nil {
			log.Println("Error reading file", file.Name(), fErr)
			continue
		}
		content := string(fileBytes)
		db.Exec(content)
		db.Create(database.Seeding{
			Name:      file.Name(),
			CreatedAt: time.Now(),
		})
		log.Printf("File %s was seeded\n", file.Name())
		count++
	}
	log.Printf("Seeding complete. Total files: %d\n", count)
}

func seedingExists(seeding []database.Seeding, seedName string) bool {
	for _, seed := range seeding {
		if seed.Name == seedName {
			return true
		}
	}
	return false
}
