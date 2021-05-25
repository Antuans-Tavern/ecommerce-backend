package seeder

import (
	"time"

	"github.com/fatih/color"
	"gorm.io/gorm"
)

type seeder func(db *gorm.DB)

func Run(db *gorm.DB) {
	start := time.Now()

	call(
		db,
		userSeeder,
		categorySeeder,
		productSeeder,
		imageSeeder,
	)

	color.Green("\n\nTook %v", time.Since(start))
}

func call(db *gorm.DB, seeders ...seeder) {
	for _, seeder := range seeders {
		seeder(db)
	}
}
