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
	)

	color.Green("\n\nTook %vs", time.Since(start))
}

func call(db *gorm.DB, seeders ...seeder) {
	for _, seeder := range seeders {
		seeder(db)
	}
}
