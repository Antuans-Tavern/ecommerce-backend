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

func Translator(locale string) ut.Translator {
	trans, _ := translator.GetTranslator(locale)
	return trans
}

func Trans(locale, msg string, params ...string) string {
	trans, _ := translator.GetTranslator(locale)
	message, _ := trans.T(msg, params...)

	return message
}
