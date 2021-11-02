package service

import (
	"github.com/SundarBank/domain"
	"github.com/SundarBank/errs"
)

type CustomerService interface{
	GetAllCustomer() ([]domain.Customer,*errs.AppError)
	ById(string) (*domain.Customer,*errs.AppError)
}

type DefaultCustomerService struct{
	repo domain.CustomerRepository
}

func (d DefaultCustomerService) GetAllCustomer() ([]domain.Customer,*errs.AppError){
	return d.repo.FindAll()
}
func (d DefaultCustomerService) ById(a string) (*domain.Customer,*errs.AppError){
	return d.repo.ById(a)
}


func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService{
	return DefaultCustomerService{repo:repository }
}
