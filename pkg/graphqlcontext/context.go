package graphqlcontext

import (
	"context"
	"net/http"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/spf13/viper"
)

type GraphqlContext struct {
	UserAgent  string
	Translator ut.Translator
}

func Middleware(next http.Handler) http.Handler {
	en := en.New()
	uni := ut.New(en, en)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		trans, _ := uni.GetTranslator(viper.GetString("lang"))
		ctx := context.WithValue(r.Context(), "contextx", GraphqlContext{
			UserAgent:  r.UserAgent(),
			Translator: trans,
		})

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
