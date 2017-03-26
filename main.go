package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync/atomic"
)

var healthy = int32(1)

func main() {

	http.HandleFunc("/healthy", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			h := atomic.LoadInt32(&healthy)
			if h == 0 {
				log.Printf("GET /healthy returns 500. (healthy: %v)", h)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			log.Printf("GET /healthy returns 204. (healthy: %v)", h)
			w.WriteHeader(http.StatusNoContent)
			return
		case http.MethodPost:
			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				log.Printf("GET /healthy returns 400: %v", err)
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprint(w, "Expected a POST body")
				return
			}

			h, err := strconv.Atoi(string(body))
			if err != nil {
				log.Printf("GET /healthy returns 400: %v", err)
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprint(w, "The body should contain an integer")
				return
			}

			log.Printf("GET /healthy returns 204. (healthy: %v)", healthy)
			atomic.StoreInt32(&healthy, int32(h))
		default:
			log.Printf("%v /healthy not allowed", r.Method)
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
