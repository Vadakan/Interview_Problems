package domain

// CustomerRepositoryStub This stub will have mock data to test our application behaviour
type CustomerRepositoryStub struct{
	customers []Customer
}

// FindAll adapter implemented the CustomerRepository interface and connected with it
func (s CustomerRepositoryStub) FindAll() ([]Customer,error){
	return s.customers,nil
}

// NewCustomerRepositoryStub Helper function to get CustomerRepositoryStub
func NewCustomerRepositoryStub() CustomerRepositoryStub {

	TotalCustomers := CustomerRepositoryStub{
		customers: []Customer{
			{"1001", "Sundar", "Theni", "500100", "31-05-1994", "Active"},
			{"1002", "Raji", "Theni", "500100", "25-01-1993", "Active"},
		},
	}
	return TotalCustomers
}
