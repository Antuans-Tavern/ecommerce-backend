package rule

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func DBUnique(db *gorm.DB) func(validator.FieldLevel) bool {
	return func(fl validator.FieldLevel) bool {
		params := strings.Split(fl.Param(), ";")

		var count int64

		tx := db.Table(params[0])

		if len(params) == 2 {
			tx = tx.Where(fmt.Sprintf("%s = '%s'", params[1], fl.Field().String()))
		} else {
			tx = tx.Where(fmt.Sprintf("%s = '%s'", strings.ToLower(fl.FieldName()), fl.Field().String()))
		}

		tx.Count(&count)
		return count == 0
	}
}
