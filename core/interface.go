package core

import (
	"net/http"
)

type IRequestHandler interface {
	Handle(r *http.Request, rc *RequestContext) ([]byte, error)
}
