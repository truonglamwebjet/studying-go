package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/nicholasjackson/building-microservices-youtube/product-api/handlers"
)

func main() {
	// http.HandleFunc expects a function => register handlers for requests in http servers
	lesson := 2

	switch lesson {
	case 1:
		http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
			log.Println("Hello World")
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
			http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
				log.Println("Goodbye World")
			})

			http.ListenAndServe(":9090", nil)
		})
	case 2:
		l := log.New(os.Stdout, "product-api", log.LstdFlags)

		// create hello handler object
		hh := handlers.NewHello(l)

		// create goodbye handler object
		gh := handlers.NewGoodbye(l)

		sm := http.NewServeMux()
		sm.Handle("/", hh)
		sm.Handle("/goodbye", gh)

		// manually create a server
		s := &http.Server{
			Addr:    ":9090",
			Handler: sm,
			// useful when a client call multiple requests within the same micro-services connecting to each other. Maintaing the connection between each other
			IdleTimeout:  120 * time.Second,
			ReadTimeout:  1 * time.Second,
			WriteTimeout: 1 * time.Second,
		}

		go func() {
			err := s.ListenAndServe()
			if err != nil {
				l.Fatal(err)
			}
		}()

		// create a channel
		sigChan := make(chan os.Signal)
		signal.Notify(sigChan, os.Interrupt)
		signal.Notify(sigChan, os.Kill)

		// reading from a channel
		sig := <-sigChan
		l.Println("Recieved terminate, graceful shutdown", sig)

		// wait till the requeset completed it will stop the connect/ wait for everyone to finish. Forcefully shut down if more than 30s
		tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
		s.Shutdown(tc)
	}

}
