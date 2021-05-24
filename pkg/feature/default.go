package feature

var features = map[string]bool{
	"default": true,
}

func HasFeature(key string) bool {
	_, ok := features[key]

	return ok
}
