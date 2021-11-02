package service

import (
	"github.com/SundarBank/domain"
	"github.com/SundarBank/dto"
	"github.com/SundarBank/errs"
)

type CustomerService interface{
	GetAllCustomer() ([]domain.Customer,*errs.AppError)
	ById(string) (*dto.CustomerResponse,*errs.AppError)
}

type DefaultCustomerService struct{
	repo domain.CustomerRepository
}

func (d DefaultCustomerService) GetAllCustomer() ([]domain.Customer,*errs.AppError){
	return d.repo.FindAll()
}
func (d DefaultCustomerService) ById(a string) (*dto.CustomerResponse,*errs.AppError){
	c,err :=  d.repo.ById(a)
	if err != nil{
		return nil,err
	}
	response := c.Todto()
	//instead of taking from below introduced new method to get from server
	/*
	response := dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateOfBirth: c.DateOfBirth,
		Status:      c.Status,
	}
	*/

	return &response,nil
}


func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService{
	return DefaultCustomerService{repo:repository }
}
