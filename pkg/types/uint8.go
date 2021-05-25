package types

import (
	"fmt"
	"io"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalUInt8(i uint32) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		io.WriteString(w, fmt.Sprintf("%d", i))
	})
}

func UnmarshalUInt8(v interface{}) (uint8, error) {
	switch v := v.(type) {
	case int:
		return uint8(v), nil
	default:
		return 0, fmt.Errorf("%T is not an uint32", v)
	}
}
