package config

import "time"

type ServiceConfig struct {
	Routes map[string]Routes `yaml:"routes"`
}

type Routes []Route

type Route struct {
	Method   HTTPMethod `yaml:"method"`
	Response Response   `yaml:"response"`
}

type Response struct {
	Status int    `yaml:"status"`
	Body   Body   `yaml:"body,omitempty"`
	Delay  *Delay `yaml:"delay,omitempty"` // Time in nanoseconds
}

type Body any
type HTTPMethod string
type Delay time.Duration

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
