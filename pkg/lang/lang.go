package lang

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/es"
	ut "github.com/go-playground/universal-translator"
)

var translator *ut.UniversalTranslator

func init() {
	es := es.New()
	en := en.New()

	translator = ut.New(es, en)

	setUpEs()
	setUpEn()
}

// Translator returns the instance of translator for a given locale.
// If the given locale doesn't exist, a fallback locale will be used
func Translator(locale string) ut.Translator {
	trans, _ := translator.GetTranslator(locale)
	return trans
}

// Trans returns a translated message for a given locale.
// If the locale doesn't exist, a fallback locale will be used
func Trans(locale, msg string, params ...string) string {
	trans, _ := translator.GetTranslator(locale)
	message, _ := trans.T(msg, params...)

	return message
}
