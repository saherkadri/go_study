package rest_api

import (
	"net/http"
)

func CreateEndpoints() {
	http.HandleFunc("/echo", Echo)
	http.HandleFunc("/invert", Invert)
	http.HandleFunc("/flatten", Flatten)
	http.HandleFunc("/sum", Sum)
	http.HandleFunc("/multiply", Multiply)
}
