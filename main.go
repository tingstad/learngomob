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
	// salsap siteconfig: curl -X GET http://services.snap0.api.no/api/config/sites/salsap -H 'Accept: application/json'
	// ben sier:
	//  - GET
	//  - POST
	//
	//  - perhaps when you post we call, send back data from external service
	//  (like siteconfig). Something JSON?

}

// retro etter mobbing runde 1.
//
// folk:
//
//  1. Vi lærte litt
//  2. Gikk sakte, kanskje fint
//  3. Fire minutter var korte iterasjoner
//  4. Vi tror dette er en effektiv måte å lære Go på
//
// ben:
//
//  1. Hva med eksisterende kode? Hva finnes? Konsistens.
