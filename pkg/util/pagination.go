package util

import (
	"github.com/go-playground/validator/v10"
)

func ValidatePaginationParams(pagination int, page int) (int, int) {
	params := map[string]interface{}{
		"pagination": pagination,
		"page":       page,
	}

	rules := map[string]interface{}{
		"pagination": "required,gte=1",
		"page":       "required,gte=1",
	}

	v := validator.New()
	errs := v.ValidateMap(params, rules)

	for key := range errs {
		params[key] = 1
	}

	return params["pagination"].(int), params["page"].(int)
}
