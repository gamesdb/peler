package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	PORT := ":42069"
	http.Handle("/", http.FileServer(http.Dir("gorokuichi")))
	fmt.Println("Listening on", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
