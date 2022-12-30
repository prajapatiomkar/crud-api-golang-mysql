package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/prajapatiomkar/crud-api-golang-mysql/config"
)

var DB *sql.DB

func Connect() error {
	var err error
	DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", config.Config("DB_USER"), config.Config("DB_PASSWORD"), config.Config("DB_NAME")))

	if err != nil {
		return err
	}
	if err = DB.Ping(); err != nil {
		return err
	}
	CreateEmployeeTable()
	fmt.Println("Connection Opened to Database")
	return nil
}

func CreateEmployeeTable() {
	DB.Query(`CREATE TABLE IF NOT EXISTS employees(
		id INT(4) PRIMARY KEY,
		name CHAR(255) NOT NULL,
		salary FLOAT(10,2) NOT NULL,
		age int(3) NOT NULL
	)`)
}
