package main

import "fmt"
import "net/http"

func main() {
	resp, _ := http.Get("http://example.com/")
	fmt.Println(resp)

}
