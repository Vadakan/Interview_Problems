package app

import (
	"encoding/json"
	"github.com/SundarBank/service"
	"github.com/gorilla/mux"
	"net/http"
)

type ClientCustomerHandler struct{
	service service.CustomerService
}

func(ch *ClientCustomerHandler) GetAllCustomer(w http.ResponseWriter, r *http.Request){
    customers, err := ch.service.GetAllCustomer()

	if err != nil{

		WriteResponse(w,err.Code,err.AsMessage())

	}else {
		WriteResponse(w,http.StatusOK,customers)
	      }
}

func(ch *ClientCustomerHandler) GetById(w http.ResponseWriter, r *http.Request){
	URL := mux.Vars(r)
	IndividualCustomer, err := ch.service.ById(URL["id"])

	if err != nil{
		WriteResponse(w,err.Code,err.AsMessage())
		}else {
		WriteResponse(w,http.StatusOK,IndividualCustomer)
	}
}

func WriteResponse(w http.ResponseWriter,code int,data interface{}){
	//Always in below order - first set the 'Content-Type' and then write the header and then encode it
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}



