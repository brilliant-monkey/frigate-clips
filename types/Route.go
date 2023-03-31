package types

import "net/http"

type Route struct {
	Method      string
	Path        string
	HandlerFunc http.Handler
	Name        string
}
