package seeder

import (
	"fmt"
	"time"

	"github.com/Antuans-Tavern/ecommerce-backend/pkg/database/model"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/fatih/color"
	"gorm.io/gorm"
)

func ProductSeeder(db *gorm.DB) {
	color.Blue("Seeding Products...")
	start := time.Now()

	products := []*model.Product{}

	categories := []*model.Category{}

	db.Find(&categories)

	for _, category := range categories {
		for i := 0; i < 10; i++ {
			products = append(products, &model.Product{
				Name:        fmt.Sprintf("%s %s", gofakeit.HipsterSentence(3), gofakeit.Word()),
				Description: gofakeit.HipsterParagraph(1, 3, 5, " "),
				CategoryID:  category.ID,
			})
		}
	}

	db.CreateInBatches(products, 10)

	color.Green("Seeded %d Products. (%v)", len(products), time.Since(start))
}
