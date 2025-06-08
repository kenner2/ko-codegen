package properties

import "github.com/santhosh-tekuri/jsonschema/v5"

type Identifier interface {
	GetType(name string, property jsonschema.Schema, isOptional bool) (string, error)
}
