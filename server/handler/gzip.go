package handler

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
	"sync"
)

var gzipWriterPool = &sync.Pool{New: func() interface{} { return gzip.NewWriter(nil) }}

type GzipHandler struct {
	InnerHandler http.Handler
}

type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func (this *GzipHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		this.InnerHandler.ServeHTTP(w, r)
		return
	}
	w.Header().Set("Content-Encoding", "gzip")
	gz := gzipWriterPool.Get().(*gzip.Writer) //this cast is safe
	defer gzipWriterPool.Put(gz)
	gz.Reset(w)
	defer gz.Close()
	gzr := gzipResponseWriter{Writer: gz, ResponseWriter: w}
	this.InnerHandler.ServeHTTP(gzr, r)
}
