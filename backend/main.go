package main

import (
	"bmo-ipmi/pkg/handlers"
	"log"
	"net/http"
)

func main() {
	handlers.InitSerial()

	http.Handle("/api/power-on", http.HandlerFunc(handlers.PowerOnHandler))
	http.Handle("/api/power-off", http.HandlerFunc(handlers.PowerOffHandler))
	http.Handle("/api/reset", http.HandlerFunc(handlers.ResetHandler))

	log.Println("Listening on http://localhost:8080")
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
