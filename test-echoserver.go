package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		b, er := ioutil.ReadAll(r.Body)
		if er != nil {
			log.Fatalln("Could not read body")
		}
		fmt.Fprintf(w, "Path %q\n", r.URL.Path)
		fmt.Fprintf(w, "Body %q\n", string(b))
		log.Printf("Path %q\n", r.URL.Path)
		log.Printf("Body %q\n", string(b))
		log.Printf("Header %v", r.Header)
	})
	http.ListenAndServe(":80", nil)
}
