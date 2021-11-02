package domain

import (
	"database/sql"
	"github.com/SundarBank/errs"
	"github.com/SundarBank/logger"
	_ "github.com/go-sql-driver/mysql"
	"time"
)


type CustomerRepositoryDB struct{
	client *sql.DB
}

func(s CustomerRepositoryDB) FindAll() ([]Customer,*errs.AppError){

	findAllSQL := "Select customer_id,name,date_of_birth,city,zipcode,status from customers"

	rows,err := s.client.Query(findAllSQL)

	if err != nil{
		logger.Error("Error while querying the customer table")
		panic(err)
	}

	customer := Customer{}
	customers := make([]Customer,0)

	for rows.Next(){
		err := rows.Scan(&customer.Id,&customer.Name,&customer.DateOfBirth,&customer.City,&customer.Zipcode,&customer.Status)
		if err != nil{
			logger.Error("Error while scanning the customer table")
			return nil,errs.NewUnexpectedDatabaseError("Error while scanning the Customer table")
		}
		customers = append(customers,customer)
	}
	return customers,nil
}

func(s CustomerRepositoryDB) ById(a string) (*Customer, *errs.AppError){

	findcustSQL := "Select customer_id,name,date_of_birth,city,zipcode,status from customers where customer_id = ?"

	row := s.client.QueryRow(findcustSQL,a)

	customer := Customer{}

	err := row.Scan(&customer.Id,&customer.Name,&customer.DateOfBirth,&customer.City,&customer.Zipcode,&customer.Status)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("No Rows found in Customer Table")
		} else {
			logger.Error("error while scanning the customer table")
			return nil, errs.NewUnexpectedDatabaseError("Unexpected Table error.might be configuration issue")
		}
	}
	return &customer,nil
}

func NewCustomerRepositoryDB () CustomerRepositoryDB{

	db,err := sql.Open("mysql","root:password@tcp(localhost:3306)/banking")

	if err != nil{
		logger.Error("Panicking due to SQL open error")
		panic(err)
	}

	db.SetConnMaxLifetime(time.Second*3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	 return CustomerRepositoryDB{db}
}
