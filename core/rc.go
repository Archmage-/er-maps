package core

import (
	"errors"
	"fmt"
	"net/http"
)

func NewRequestContext(r *http.Request) *RequestContext {
	return &RequestContext{}
}

func DontPanicRc(e *error, rcp **RequestContext) {
	if ex := recover(); ex != nil {
		switch exception := ex.(type) {
		case error:
			*e = exception
			break
		case string:
			*e = errors.New(exception)
			break
		default:
			*e = fmt.Errorf("%+#v\n", exception)
			break
		}
		//		var rc *RequestContext
		//		if rcp != nil {
		//			rc = *rcp
		//		}
		//TODO:log
	}
	return
}
