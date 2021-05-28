package seeder

import (
	"time"

	"github.com/Antuans-Tavern/ecommerce-backend/pkg/database/model"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/fatih/color"
	"gorm.io/gorm"
)

func CategorySeeder(db *gorm.DB) {
	color.Blue("Seeding Categories...")
	start := time.Now()

	categories := seedCategories(10, nil)

	db.CreateInBatches(&categories, 10)
	color.Green("Seeded %d categories. (%v)", len(categories), time.Since(start))

	subCategories := []*model.Category{}

	for _, category := range categories {
		subCategories = append(subCategories, seedCategories(5, &category.ID)...)
	}

	db.CreateInBatches(&subCategories, 10)
	color.Green("Seeded %d subcategories. (%v)", len(subCategories), time.Since(start))
}

func seedCategories(count int, parent *uint) []*model.Category {
	categories := []*model.Category{}

	for i := 0; i < count; i++ {
		categories = append(categories, &model.Category{
			Name:     gofakeit.HipsterSentence(3),
			ParentID: parent,
		})
	}

	return categories
}
