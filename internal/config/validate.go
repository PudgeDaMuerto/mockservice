package config

import (
	"fmt"
	"time"

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

func (d *Delay) UnmarshalYAML(n ast.Node) error {
	delay := n.String()
	duration, err := time.ParseDuration(delay)

	if err != nil {
		return &yaml.SyntaxError{
			Message: err.Error(),
			Token:   n.GetToken(),
		}
	}

	*d = Delay(duration)
	return nil

}
