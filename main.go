package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/CreamyMilk/koacurrency/router"
)

func main() {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = ":10001"
	}

	mux := http.NewServeMux()
	router.SetupRoutes(mux)
	fmt.Printf("Listening on port %s\n", PORT)
	http.ListenAndServe(PORT, router.Logger(mux))
}
