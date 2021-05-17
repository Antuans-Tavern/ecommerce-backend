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

	for _, cat := range gofakeit.Categories() {
		categories = append(categories, &model.Category{
			Name: cat[0],
		})
	}

	db.CreateInBatches(categories, 5)

	color.Green("Seeded %d Users. (%vs)", len(categories), time.Since(start).Seconds())
}
