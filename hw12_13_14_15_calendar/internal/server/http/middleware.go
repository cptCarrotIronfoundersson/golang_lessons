package internalhttp

import (
	"errors"
	"fmt"
	"net/http"
)

type statusWriter struct {
	http.ResponseWriter
	status int
}
type middleware func(logger Logger, next http.Handler) http.Handler

func middlewareChainApply(logger Logger, next http.Handler, m []middleware) http.Handler {
	chainedHandler := next
	for _, mid := range m {
		chainedHandler = mid(logger, chainedHandler)
	}
	return chainedHandler
}

func (w *statusWriter) WriteHeader(status int) {
	w.status = status
	w.ResponseWriter.WriteHeader(status)
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

func EnsureAppJSONMiddleware(logger Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			fmt.Println(r.Header.Get("Content-Type"))
			w.WriteHeader(http.StatusBadRequest)
			next.ServeHTTP(w, r)
			errmsg := fmt.Sprintf("StatusCode: %d - there is no application/json Content-type",
				http.StatusBadRequest)
			w.Write([]byte(errmsg))
			logger.Error(errors.New(errmsg))
			return
		}
		next.ServeHTTP(w, r)
	})
}
