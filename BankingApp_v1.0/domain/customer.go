package domain

// Customer This is business domain

type Customer struct{

	Id          string
	Name        string
	City        string
	Zipcode     string
	DateOfBirth string
	Status      string

}

//CustomerRepository This is primary port interface
type CustomerRepository interface{
	FindAll()([]Customer,error)
}

