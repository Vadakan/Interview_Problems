package app

import (
	"fmt"
	"github.com/SundarBank/domain"
	"github.com/SundarBank/service"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func Sanitycheck(){

	if os.Getenv("SERVER_ADDRESS") == ""||
		os.Getenv("SERVER_PORT") == "" ||
		os.Getenv("DB_USER") == "" ||
		os.Getenv("DB_PASSWORD") == "" ||
		os.Getenv("DB_PROTOCOL") == "" ||
		os.Getenv("DB_ADDRESS") == "" ||
		os.Getenv("DB_PORT") == "" ||
		os.Getenv("DB_NAME") == ""{
		log.Fatal("Environment Variables are not defined...")
	}
}

func Start(){

	Sanitycheck()

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

	 address := os.Getenv("SERVER_ADDRESS")
	 port := os.Getenv("SERVER_PORT")
	 //http.ListenAndServe(":8080",router)
	 http.ListenAndServe(fmt.Sprintf("%s:%s",address,port),router)

}