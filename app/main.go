package main

import (
	"fmt"
	"github.com/rizaramadan/go102"
	"github.com/rizaramadan/go102/webs"
	"html"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "this is home, %q \n",
		html.EscapeString(r.URL.Path))
	if err != nil {
		log.Println(err)
	}
}

func main() {
	route := webs.NewRoute("/", home)
	route.AddRoute()

	log.Println("listen localhost 8989")
	log.Fatal(http.ListenAndServe(go102.Port, nil))
}
