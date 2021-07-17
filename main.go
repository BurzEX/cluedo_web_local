package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/cluedo/", http.StripPrefix("/cluedo/", fs))

	fmt.Println("localhost:8080/cluedo")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

// func handler(w http.ResponseWriter, r *http.Request) {
// 	// fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
// 	http.ServeFile(w, r, "static/index.html")
// }
