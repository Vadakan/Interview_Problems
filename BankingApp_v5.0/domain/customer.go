package domain

import (
	"github.com/SundarBank/dto"
	"github.com/SundarBank/errs"
)

type Customer struct{
	Id          string  `db:"customer_id"`
	Name        string  `db:"name"`
	City        string  `db:"city"`
	Zipcode     string  `db:"zipcode"`
	DateOfBirth string  `db:"date_of_birth"`
	Status      string  `db:"status"`
}

func(c Customer) StatusAsText() string{
	statusAstext := "Active"
	if c.Status == "0"{
		statusAstext = "InActive"
	}
	return statusAstext
}

func (c Customer) Todto() dto.CustomerResponse {

	return dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.StatusAsText(),
	}
}

type CustomerRepository interface {
	FindAll()([]Customer,*errs.AppError)
	ById(string) (*Customer,*errs.AppError)
}
