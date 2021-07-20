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
	fmt.Printf("Listening on port %s\n", PORT)
	router.SetupRoutes(mux)
	http.ListenAndServe(PORT, mux)
}
