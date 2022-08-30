package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Hello WOrld")
		d, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "oops", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(rw, "Hello %s\n", d)
	})
	http.HandleFunc("/goodbye", func(http.ResponseWriter, *http.Request) {
		log.Println("goodbye WOrld")
	})

	http.ListenAndServe(":9090", nil)
}
