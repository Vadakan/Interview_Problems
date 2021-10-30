package app

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Start() {

	//create a router
	router := mux.NewRouter()

	//API end point to fan-out the request
	router.HandleFunc("/customer", GetAllCustomers).Methods(http.MethodGet)
	router.HandleFunc("/greet", HelloFunc).Methods(http.MethodGet)
	router.HandleFunc("/customer/{id:[0-9]+}",CheckIntType).Methods(http.MethodGet)
	router.HandleFunc("/customer/{name:[A-Z]+[a-z]+[0-9]+}",CheckStringType).Methods(http.MethodGet)

	//starting the server
	http.ListenAndServe(":8080", router)

}