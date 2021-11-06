package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main(){

	router := mux.NewRouter()
    router.HandleFunc("/Home",getHome).Methods(http.MethodGet).Name("GettingHome")
	router.HandleFunc("/Login",getLogin).Methods(http.MethodPost).Name("GettingHome")

	//middleware function
	router.Use(authorizationheader())

	http.ListenAndServe(":8080",router)

}

func authorizationheader() func(http.Handler) http.Handler{
	fmt.Println("Inside middleware function ")

    var Token_Auth bool
	Token_Auth = true
	return func(next http.Handler) http.Handler{
		var Hfunc http.HandlerFunc
		Hfunc = func(w http.ResponseWriter, r *http.Request) {
			if Token_Auth {
				RouteName := mux.CurrentRoute(r).GetName()
				fmt.Println("RouteName inside middleware func :",RouteName )
				next.ServeHTTP(w,r)
			}else{
				Token_Auth = false
				fmt.Println("Unauthorized Token ")
			}
		}

		return Hfunc
	}

}

func getHome(w http.ResponseWriter, r *http.Request){
	//below Current route will help in getting all the details about the route
	RouteName := mux.CurrentRoute(r).GetName()
	fmt.Println("RouteName :",RouteName )

	sliceOfMethods, _  := mux.CurrentRoute(r).GetMethods()
	fmt.Println("Methods :", sliceOfMethods)

	//below method get route variables if we use and Regex variables
	vars := mux.Vars(r)
	fmt.Println("Route Variable :",vars)

	authHeader := r.Header.Get("Authorization")
	fmt.Println("Bearer Token :",authHeader)
	fmt.Printf("%T\n",authHeader)

}

func getLogin(w http.ResponseWriter, r *http.Request){
	//below Current route will help in getting all the details about the route
	RouteName := mux.CurrentRoute(r).GetName()
	fmt.Println("RouteName :",RouteName )

	sliceOfMethods, _  := mux.CurrentRoute(r).GetMethods()
	fmt.Println("Methods :", sliceOfMethods)

	//below method get route variables if we use and Regex variables
	vars := mux.Vars(r)
	fmt.Println("Route Variable :",vars)

	authHeader := r.Header.Get("Authorization")
	fmt.Println("Bearer Token :",authHeader)
	fmt.Printf("%T\n",authHeader)


}
