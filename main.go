package main

import (
	"github.com/antoniocarlosmjr/routes"
	"net/http"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":1010", nil)
}
