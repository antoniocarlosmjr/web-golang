package routes

import (
	"github.com/antoniocarlosmjr/controllers"
	"net/http"
)

func CarregaRotas() {
	http.ListenAndServe(":1010", nil)
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/new", controllers.New)
}
