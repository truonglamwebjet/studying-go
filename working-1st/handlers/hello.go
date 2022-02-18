package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	// Dependency injection:
	// write good fast unit-tests.
	// Rely on interface instead of object. We can create fake object to test( mock out test object)
	return &Hello{l}
}

// signature for this handle => satisfy the http handler
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Hello World")
	d, err := ioutil.ReadAll(r.Body)

	if err != nil {
		useErrorFunc := 1
		if useErrorFunc == 1 {
			//similar to the option down below but only need 1 line of code
			http.Error(rw, "Ooops", http.StatusBadRequest)
		} else {
			// allow to write back the error to header
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("Ooops"))
		}
		return
	}

	fmt.Fprintf(rw, "Hello %s", d)
}
