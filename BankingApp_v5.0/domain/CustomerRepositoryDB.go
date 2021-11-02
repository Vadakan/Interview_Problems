package domain

import (
	"database/sql"
	"fmt"
	"github.com/SundarBank/errs"
	"github.com/SundarBank/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"os"
	"time"
)


type CustomerRepositoryDB struct{
	client *sqlx.DB
}

func(s CustomerRepositoryDB) FindAll() ([]Customer,*errs.AppError){

	findAllSQL := "Select customer_id,name,date_of_birth,city,zipcode,status from customers"

	customers := make([]Customer,0)

    err := s.client.Select(&customers,findAllSQL)

	if err != nil{
		logger.Error("Error while querying the customer table " + err.Error())
		panic(err)
	}

	return customers,nil
}

func(s CustomerRepositoryDB) ById(a string) (*Customer, *errs.AppError){

	findcustSQL := "Select customer_id,name,date_of_birth,city,zipcode,status from customers where customer_id = ?"

	customer := Customer{}
	err := s.client.Get(&customer,findcustSQL,a)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("No Rows found in Customer Table " + err.Error())
		} else {
			logger.Error("error while scanning the customer table " + err.Error())
			return nil, errs.NewUnexpectedDatabaseError("Unexpected Table error.might be configuration issue " + err.Error())
		}
	}
	return &customer,nil
}

func NewCustomerRepositoryDB () CustomerRepositoryDB{

	Username := os.Getenv("DB_USER")
	Password := os.Getenv("DB_PASSWORD")
	Protocol := os.Getenv("DB_PROTOCOL")
	Network  := os.Getenv("DB_ADDRESS")
	Port := os.Getenv("DB_PORT")
	DatabaseName := os.Getenv("DB_NAME")

	SQL_conn_string := fmt.Sprintf("%s:%s@%s(%s:%s)/%s",Username,Password,Protocol,Network,Port,DatabaseName)

	db,err := sqlx.Open("mysql",SQL_conn_string)

	if err != nil{
		logger.Error("Panicking due to SQL open error" + err.Error())
		panic(err)
	}

	db.SetConnMaxLifetime(time.Second*3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	 return CustomerRepositoryDB{db}
}
