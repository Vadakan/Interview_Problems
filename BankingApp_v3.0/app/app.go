package app

import (
	"github.com/SundarBank/domain"
	"github.com/SundarBank/service"
	"github.com/gorilla/mux"
	"net/http"
)

func Start(){

	 router := mux.NewRouter()
     /*
	 ch := ClientCustomerHandler{
		 service: service.NewCustomerService(domain.NewCustomerRepository()),
	 }
     */

	 ch := ClientCustomerHandler{
		 service: service.NewCustomerService(domain.NewCustomerRepositoryDB()),
	 }



	router.HandleFunc("/customer",ch.GetAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customer/{id:[0-9]+}",ch.GetById).Methods(http.MethodGet)

	 http.ListenAndServe(":8080",router)

}