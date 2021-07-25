package middlewares

import (
	"log"
	"net/http"
	"time"
)

type Logger struct {
	handler http.Handler
}

func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handlerToWrap}
}

func (logger *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	logger.handler.ServeHTTP(w, r)

	log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
}
