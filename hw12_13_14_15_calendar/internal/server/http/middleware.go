package internalhttp

import (
	"fmt"
	"net/http"
)

type statusWriter struct {
	http.ResponseWriter
	status int
	length int
}

func (w *statusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
}

func (w *statusWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.status = 200
	}
	n, err := w.ResponseWriter.Write(b)
	w.length += n
	return n, err
}

func loggingMiddleware(logger Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sw := statusWriter{ResponseWriter: w}
		next.ServeHTTP(w, r)
		logger.Info(
			fmt.Sprintf("ClientAddr: %s \n Method: %s\n URL: %s\nHttpProtocol: %s\nStatusCode: %d",
				r.Host, r.Method, r.URL, r.Proto, sw.status),
		)
	})
}
