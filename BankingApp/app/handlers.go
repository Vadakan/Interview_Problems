package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type Customer struct{
	Name string  	`json:"name"    xml:"name"`
	City string     `json:"city"    xml:"city"`
	Zipcode string  `json:"zipcode" xml:"zipcode"`
}


func HelloFunc(w http.ResponseWriter,r *http.Request){

	fmt.Fprint(w,"Hello-World")

}

func GetAllCustomers(w http.ResponseWriter, r *http.Request){

	Customers := []Customer{
		{"Sundar","Madras","500300"},
		{"Raji","Bangalore","500400"},
	}

	if r.Header.Get("Content-Type") == "application/json" {

		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Customers)

	}else if r.Header.Get("Content-Type") == "text/plain"{

		w.Header().Add("Content-Type","text/plain")
		fmt.Fprint(w,Customers)

	}else{

		w.Header().Add("Content-Type","application/xml")
		xml.NewEncoder(w).Encode(Customers)

	}

}

func CheckIntType(w http.ResponseWriter, r *http.Request){
	Url := mux.Vars(r)
	v,_ := Url["id"]
	fmt.Printf("%T\n",v)
	fmt.Fprint(w,v," Integer value recieved")
}

func CheckStringType(w http.ResponseWriter, r *http.Request){
	Url := mux.Vars(r)
	v,_ := Url["id"]
	fmt.Printf("%T\n",v)
	fmt.Fprint(w,v," string value recieved")
}