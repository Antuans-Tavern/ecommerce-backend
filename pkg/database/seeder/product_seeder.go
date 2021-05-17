package seeder

import (
	"time"

	"github.com/Antuans-Tavern/ecommerce-backend/pkg/database/model"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/fatih/color"
	"gorm.io/gorm"
)

func productSeeder(db *gorm.DB) {
	color.Cyan("Seeding Products...")
	start := time.Now()

	products := []*model.Product{}

	categories := []*model.Category{}

	db.Find(&categories)

	for _, category := range categories {
		for i := 0; i < 10; i++ {
			products = append(products, &model.Product{
				Name:       gofakeit.HipsterWord(),
				CategoryID: category.ID,
			})
		}
	}

	db.CreateInBatches(products, 20)

	color.Green("Seeded %d Products. (%vs)", len(products), time.Since(start).Seconds())
}
