package domain

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Surender-Kumar-1996/sample_banking/config"
	"github.com/Surender-Kumar-1996/sample_banking/errs"
	"github.com/Surender-Kumar-1996/sample_banking/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDb) FindAll(status string) ([]Customer, *errs.AppError) {
	var findAllSql string
	var err error
	customers := make([]Customer, 0)
	if status == "" {
		findAllSql = "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		err = d.client.Select(&customers, findAllSql)
		// rows, err = d.client.Query(findAllSql)
	} else {
		findAllSql = "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		err = d.client.Select(&customers, findAllSql, status)
		// rows, err = d.client.Query(findAllSql, status)
	}
	if err != nil {
		logger.Error("Error while quering customer table " + err.Error())
		return nil, errs.NewNotFoundError("customer not found")
	}

	// if err := sqlx.StructScan(rows, &customers); err != nil {
	// 	logger.Error("Error while scanning customer table " + err.Error())
	// 	return nil, errs.NewUnexpectedError("Unexpected database error")
	// }
	// for rows.Next() {
	// 	var c Customer
	// 	err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	// 	if err != nil {
	// 		logger.Error("Error while scanning customer table " + err.Error())
	// 		return nil, errs.NewUnexpectedError("Unexpected database error")
	// 	}
	// 	customers = append(customers, c)
	// }
	return customers, nil
}

func NewCustomerRepositoryDb(conf *config.BankingConfig) CustomerRepositoryDb {
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", conf.Database.DbUser, conf.Database.DbPassword, conf.Database.DbAddr, conf.Database.DbPort, conf.Database.DbName)
	client, err := sqlx.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{
		client: client,
	}
}

func (d CustomerRepositoryDb) ById(id string) (*Customer, *errs.AppError) {
	var c Customer

	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ?"

	// rows := d.client.QueryRow(findAllSql, id)
	err := d.client.Get(&c, findAllSql, id)
	// err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("customer not found")
		}
		logger.Error("Error while scanning customer" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &c, nil

}
