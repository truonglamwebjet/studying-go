package handlers

import (
	"log"
	"net/http"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	// Dependency injection:
	// write good fast unit-tests.
	// Rely on interface instead of object. We can create fake object to test( mock out test object)
	return &Goodbye{l}
}

// signature for this handle => satisfy the http handler
func (h *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Byeee"))
}
