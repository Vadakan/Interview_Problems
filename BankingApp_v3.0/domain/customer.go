package domain

import "github.com/SundarBank/errs"

type Customer struct{
	Id          string  `db:"customer_id"`
	Name        string  `db:"name"`
	City        string  `db:"city"`
	Zipcode     string  `db:"zipcode"`
	DateOfBirth string  `db:"date_of_birth"`
	Status      string  `db:"status"`
}

type CustomerRepository interface {
	FindAll()([]Customer,*errs.AppError)
	ById(string) (*Customer,*errs.AppError)
}
