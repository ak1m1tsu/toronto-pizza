package main

import (
	"net/http"

	"github.com/romankravchuk/toronto-pizza/internal/router"
)

func main() {
	r := router.NewRouter()
	http.ListenAndServe(":3000", r)
}
