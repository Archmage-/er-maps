package handler

import (
	//	"bytes"
	"compress/gzip"
	"fmt"
	//	"github.com/Archmage-/er-maps/consts"
	"github.com/Archmage-/er-maps/core"
	"net/http"
	"strings"
	//	"time"
)

//RContextHandler creates RequestContext and delegates call to the inner request handler
type RContextHandler struct {
	InnerHandler core.IRequestHandler
}

func (this *RContextHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//	if r.Method == "OPTIONS" {
	//		w.Header().Set(consts.HEADER_ACCESSCONTROLALLOWHEADERS, "Content-Type, Authorization, Cookie, Pragma,  Content-Type")
	//		w.Header().Set(consts.HEADER_ACCESSCONTROLALLOWMETHODS, "OPTIONS, GET, POST")
	//		w.Header().Set(consts.HEADER_ACCESSCONTROLALLOWORIGIN, "*")
	//		w.Header().Set(consts.HEADER_ALLOW, "OPTIONS, GET, POST")
	//		return
	//	}

	fmt.Println("RC")
	var err error
	rc := core.NewRequestContext(r)

	var data []byte
	if this.InnerHandler != nil {

		fmt.Println("RC1")
		data, err = this.safeCall(r, rc)
	}

	fmt.Println("RC2")

	this.FormatResponse(w, r, rc, data, err)
}

func (this *RContextHandler) FormatResponse(w http.ResponseWriter, r *http.Request, rc *core.RequestContext, data []byte, err error) {
	//	start := time.Now()

	//GZip or not GZip
	if err == nil && data != nil && len(data) > 860 {
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			w.Header().Set("Content-Encoding", "gzip")
			gz := gzipWriterPool.Get().(*gzip.Writer) //this cast is safe
			defer gzipWriterPool.Put(gz)
			gz.Reset(w)
			defer gz.Close()
			w = gzipResponseWriter{Writer: gz, ResponseWriter: w}
		}
	}

	//-=AFTER THIS LINE NO RESPONSE HEADER CAN BE WRITTEN=-

	//Write data (if any)
	if err == nil {
		w.WriteHeader(http.StatusOK)
		w.Write(data)
	} else {
		//header
		w.WriteHeader(http.StatusInternalServerError)
		//data
		w.Write([]byte(err.Error()))
		if data != nil {
			w.Write(data)
		}
	}

}

func (this *RContextHandler) safeCall(r *http.Request, rc *core.RequestContext) (data []byte, e error) {
	defer core.DontPanicRc(&e, &rc)

	return this.InnerHandler.Handle(r, rc)
}
