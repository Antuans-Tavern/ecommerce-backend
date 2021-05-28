package seeder

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/fatih/color"
	"gorm.io/gorm"
)

type seeder func(db *gorm.DB)

func Run(db *gorm.DB) {
	start := time.Now()

	Call(
		db,
		UserSeeder,
		CategorySeeder,
		ProductSeeder,
		ImageSeeder,
	)

	color.Green("\n\nTook %v", time.Since(start))
}

func Call(db *gorm.DB, seeders ...seeder) {
	gofakeit.Seed(42)

	for _, seeder := range seeders {
		seeder(db)
	}
}
