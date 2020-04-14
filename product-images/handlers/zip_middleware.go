package handlers

import (
	"compress/gzip"
	"net/http"
	"strings"
)

type GzipHandler struct {
}

func (g *GzipHandler) GzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		if strings.Contains(req.Header.Get("Accept-Encoding"), "gzip") {
			// create a gzip response
			wrw := NewWrappedResponseWrite(res)
			wrw.Header().Set("Content-Encoding", "gzip")
			next.ServeHTTP(wrw, req)
			defer wrw.Flush()
			return
		}
		next.ServeHTTP(res, req)
	})
}

type WrappedResponseWriter struct {
	res http.ResponseWriter
	gw  *gzip.Writer
}

func NewWrappedResponseWrite(res http.ResponseWriter) *WrappedResponseWriter {
	gw := gzip.NewWriter(res)
	return &WrappedResponseWriter{res: res, gw: gw}
}

func (wr *WrappedResponseWriter) Header() http.Header {
	return wr.res.Header()
}

func (wr *WrappedResponseWriter) Write(d []byte) (int, error) {
	return wr.gw.Write(d)
}

func (wr *WrappedResponseWriter) WriteHeader(statusCode int) {
	wr.res.WriteHeader(statusCode)
}

func (wr *WrappedResponseWriter) Flush() {
	wr.gw.Flush()
	wr.gw.Close()
}
