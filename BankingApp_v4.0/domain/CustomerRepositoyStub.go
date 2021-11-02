package domain

import (
	"github.com/SundarBank/errs"
)

type CustomerRepositoryStub struct{
	Customers []Customer
}

func(c CustomerRepositoryStub) FindAll() ([]Customer,*errs.AppError){
	return c.Customers,nil
}

func(c CustomerRepositoryStub) ById(a string) (cust *Customer,e *errs.AppError){
	for _, IndividualCustomer := range c.Customers {
		if IndividualCustomer.Id == a{
			cust = &IndividualCustomer
			e = nil
			break
		}else{
			cust = nil
			e = errs.NewNotFoundError("Customer not available in the database")
		}
	}
	return cust,e
}

func NewCustomerRepository() CustomerRepositoryStub{
	customers := []Customer{
		{"1001","Rajesh","Theni","5000001","31-05-1994","Active"},
		{"1002","Raji","Theni","5000001","31-05-1994","Active"},
	}
	return CustomerRepositoryStub{customers}
}