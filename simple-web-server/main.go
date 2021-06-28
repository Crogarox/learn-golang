package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

var counter int
var mutex = &sync.Mutex{}

func echoString(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func incrementCounter(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	fmt.Fprintf(w, strconv.Itoa(counter))
	mutex.Unlock()
}

func main() {

	// tells the http package to handle all requests to the web root ("/") with handler.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// create a sub-slice of Path from the 1st character to the end.
		// This drops the leading "/" from the path name.
		fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
		// http.ServeFile(w, r, r.URL.Path[1:])
	})

	// tells the http package to handle all requests to the web path ("/increment") with handler.
	http.HandleFunc("/increment", incrementCounter)

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	http.HandleFunc("/echo", echoString)

	// http.Handle("/", http.FileServer(http.Dir("./static")))

	// Listen on any interface on port 8080
	// log.Fatal(http.ListenAndServe(":8080", nil))
	log.Fatal(http.ListenAndServeTLS(":443", "server.crt", "server.key", nil))

}
