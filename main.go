package main

import (
	"fmt"
	"net/http"
	"time"
)

func middlewareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/* Logic before the request */
		fmt.Println("before handler; middleware start")
		start := time.Now()
		/* handle request */
		handler.ServeHTTP(w, r)
		/* Logic after the request */
		fmt.Printf("middleware finished; %s", time.Since(start))
	})
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("This is a get request"))
	case http.MethodPost:
		w.Write([]byte("This is a post request"))
	}

}

func main() {
	http.Handle("/foo", middlewareHandler(http.HandlerFunc(fooHandler)))

	http.ListenAndServe(":5000", nil)
}
