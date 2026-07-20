package config

import (
	"fmt"

	"github.com/goccy/go-yaml"
	"github.com/goccy/go-yaml/ast"
)

func (m *HTTPMethod) UnmarshalYAML(n ast.Node) error {
	method := HTTPMethod(n.String())

	if _, ok := validMethods[method]; !ok {
		return &yaml.SyntaxError{
			Message: fmt.Sprintf(errInvalidHttpMethod, method),
			Token:   n.GetToken(),
		}
	}

	*m = method
	return nil
}
