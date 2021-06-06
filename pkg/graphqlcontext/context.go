package graphqlcontext

import (
	"context"
	"net/http"

	"github.com/spf13/viper"
)

type Context struct {
	UserAgent string
	Locale    string
}

func GetContext(ctx context.Context) Context {
	return ctx.Value("context").(Context)
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := context.WithValue(r.Context(), "context", Context{
			UserAgent: r.UserAgent(),
			Locale:    locale(r.Header.Get("Accept-Language")),
		})

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func locale(lang string) string {
	if lang == "es" || lang == "en" {
		return lang
	}

	return viper.GetString("lang")
}
