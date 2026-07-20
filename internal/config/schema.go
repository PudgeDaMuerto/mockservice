package config

type ServiceConfig struct {
	Routes map[string]Routes `yaml:"routes"`
}

type Routes []Route

type Route struct {
	Method   HTTPMethod `yaml:"method"`
	Response Response   `yaml:"response"`
}

type Response struct {
	Status int  `yaml:"status"`
	Body   Body `yaml:"body,omitempty"`
}

type Body any
type HTTPMethod string

const (
	GET    HTTPMethod = "GET"
	POST   HTTPMethod = "POST"
	PUT    HTTPMethod = "PUT"
	PATCH  HTTPMethod = "PATCH"
	DELETE HTTPMethod = "DELETE"
)

var validMethods = map[HTTPMethod]struct{}{
	GET:    {},
	POST:   {},
	PUT:    {},
	PATCH:  {},
	DELETE: {},
}
