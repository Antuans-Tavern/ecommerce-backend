package seeder

import (
	"time"

	"github.com/Antuans-Tavern/ecommerce-backend/pkg/database/model"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/fatih/color"
	"gorm.io/gorm"
)

func categorySeeder(db *gorm.DB) {
	color.Blue("Seeding Categories...")
	start := time.Now()

	categories := []*model.Category{}

	for i := 0; i < 10; i++ {
		categories = append(categories, &model.Category{
			Name: gofakeit.HipsterWord(),
		})
	}

	db.CreateInBatches(categories, 10)

	color.Green("Seeded %d categories. (%v)", len(categories), time.Since(start))
}
