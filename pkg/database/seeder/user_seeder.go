package seeder

import (
	"time"

	"github.com/Antuans-Tavern/ecommerce-backend/pkg/database/model"
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/util"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/fatih/color"
	"gorm.io/gorm"
)

func userSeeder(db *gorm.DB) {
	color.Blue("Seeding Users...")
	start := time.Now()

	users := []*model.Profile{}

	for i := 0; i < 100; i++ {
		users = append(users, &model.Profile{
			Name:     gofakeit.Name(),
			Lastname: gofakeit.LastName(),
			User: &model.User{
				Email:    gofakeit.Email(),
				Password: util.Hash("secret"),
				Status:   gofakeit.Bool(),
				Type: gofakeit.RandomInt([]int{
					0, 1,
				}),
			},
		})
	}

	users = append(users,
		&model.Profile{
			Name:     gofakeit.Name(),
			Lastname: gofakeit.LastName(),
			User: &model.User{
				Email:    "admin@testing.com",
				Password: util.Hash("secret"),
				Status:   gofakeit.Bool(),
				Type: gofakeit.RandomInt([]int{
					0, 1,
				}),
			},
		},
	)

	db.CreateInBatches(users, 100)
	color.Green("Seeded %d Users. (%v)", len(users), time.Since(start))
}
