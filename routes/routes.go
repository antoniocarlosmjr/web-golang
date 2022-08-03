package routes

import (
	"github.com/antoniocarlosmjr/controllers"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.ListenAndServe(":1010", nil)
}
