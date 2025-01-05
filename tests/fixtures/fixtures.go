package fixtures

import (
	"currency-service/models"
	"fmt"
	"log"

	"github.com/go-testfixtures/testfixtures/v3"
	"gorm.io/gorm"
)

func LoadDataFromFile[T any](db *gorm.DB, filePath string) error {
	sqlDB, err := db.DB()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB from GORM: %w", err)
	}

	fixtures, err := testfixtures.New(
		testfixtures.Database(sqlDB),
		testfixtures.Dialect("mysql"),
		testfixtures.Files(filePath),
	)
	if err != nil {
		return fmt.Errorf("error creating fixtures: %w", err)
	}

	log.Println("Loading fixture from file:", filePath)
	if err := fixtures.Load(); err != nil {
		return fmt.Errorf("error loading fixtures: %w", err)
	}

	log.Println("Fixture successfully loaded.")
	return nil
}

func LoadRatesFromFile(db *gorm.DB, filepath string) error {
	log.Println("Loading currency rates from file:", filepath)
	return LoadDataFromFile[models.Rate](db, filepath)
}
