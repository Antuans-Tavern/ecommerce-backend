package rule

import (
	"fmt"
	"strings"

	"github.com/Antuans-Tavern/ecommerce-backend/pkg/lang"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func exists(db *gorm.DB) func(validator.FieldLevel) bool {
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
		return count > 0
	}
}

func registerExistsRule(db *gorm.DB, validate *validator.Validate) {

	validate.RegisterValidation("exists", dbUnique(db))

	validate.RegisterTranslation("exists", lang.Translator("es"), func(ut ut.Translator) error {
		return ut.Add("exists", "{0} no existe.", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("exists", fe.StructField())

		return t
	})

	validate.RegisterTranslation("exists", lang.Translator("en"), func(ut ut.Translator) error {
		return ut.Add("exists", "The {0} already exists.", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("exists", fe.StructField())

		return t
	})
}
