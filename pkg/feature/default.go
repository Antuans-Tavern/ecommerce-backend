package feature

var features = map[string]bool{
	"default": true,
}

func HasFeature(key) bool {
	_, ok := features[key]

	return ok
}
