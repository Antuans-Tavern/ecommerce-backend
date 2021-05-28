package seeder

import (
	"time"

	"github.com/Antuans-Tavern/ecommerce-backend/pkg/database/model"
	"github.com/fatih/color"
	"gorm.io/gorm"
)

func ImageSeeder(db *gorm.DB) {
	color.Blue("Seeding images...")
	start := time.Now()

	images := []model.Image{}

	images = append(images, productImages(db)...)

	db.CreateInBatches(&images, 10)
	color.Green("Seeded %d images. (%v)", len(images), time.Since(start))
}

func productImages(db *gorm.DB) []model.Image {
	products := []model.Product{}
	db.Find(&products)

	images := []model.Image{}
	for _, product := range products {
		images = append(images, model.Image{
			Path:          "/200",
			ImageableType: "products",
			ImageableID:   product.ID,
			Disk:          model.IMAGE_DISK_SEED,
			Type:          model.IMAGE_TYPE_PRODUCT_MAIN,
		})
	}

	return images
}
