package middleware

import "net/http"

type ResponseRecorder struct {
	writer http.ResponseWriter
	status int
}

func (r *ResponseRecorder) Header() http.Header {
	return r.writer.Header()
}

func (r *ResponseRecorder) Write(b []byte) (int, error) {
	return r.writer.Write(b)
}

func (r *ResponseRecorder) WriteHeader(statusCode int) {
	r.status = statusCode
	r.writer.WriteHeader(statusCode)
}