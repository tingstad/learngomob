package main

import "fmt"
import "net/http"
import "html"

func fooHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/foo", fooHandler)

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	//	log.Fatal(http.ListenAndServe(":8080", nil))
	http.ListenAndServe(":8080", nil)

	// ben sier:
	//  - GET
	//  - POST
	//
	//  - perhaps when you post we call, send back data from external service
	//  (like siteconfig). Something JSON?
}
