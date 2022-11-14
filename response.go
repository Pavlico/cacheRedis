package cacheRedis

import (
	"bytes"
	"net/http"
)

//To get response body from responsewriter
type ResponseWriterWrapper struct {
	w          *http.ResponseWriter
	body       *bytes.Buffer
	statusCode *int
}

func NewResponseWriterWrapper(w http.ResponseWriter) ResponseWriterWrapper {
	var buf bytes.Buffer
	var statusCode int = http.StatusOK
	return ResponseWriterWrapper{
		w:          &w,
		body:       &buf,
		statusCode: &statusCode,
	}
}

func (rww ResponseWriterWrapper) Write(buf []byte) (int, error) {
	rww.body.Write(buf)
	return (*rww.w).Write(buf)
}

func (rww ResponseWriterWrapper) Header() http.Header {
	return (*rww.w).Header()

}

func (rww ResponseWriterWrapper) WriteHeader(statusCode int) {
	(*rww.statusCode) = statusCode
	(*rww.w).WriteHeader(statusCode)
}

func (rww ResponseWriterWrapper) GetBody() string {
	var buf bytes.Buffer
	buf.WriteString(rww.body.String())
	return buf.String()
}
