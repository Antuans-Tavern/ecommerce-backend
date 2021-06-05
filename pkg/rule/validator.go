package rule

import (
	"github.com/Antuans-Tavern/ecommerce-backend/pkg/lang"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	es_translations "github.com/go-playground/validator/v10/translations/es"
	"gorm.io/gorm"
)

func SetUpValidator(db *gorm.DB) *validator.Validate {
	validate := validator.New()

	registerDBUniqueRule(db, validate)

	es_translations.RegisterDefaultTranslations(validate, lang.Translator("es"))
	en_translations.RegisterDefaultTranslations(validate, lang.Translator("en"))

	return validate
}
