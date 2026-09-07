package config

import (
	config "AuthInGo/config/env"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func SetupDB() (*sql.DB,error) {
	cfg:=mysql.NewConfig()

	cfg.User=config.GetString("DB_USER","root")
	cfg.Passwd=config.GetString("DB_PASSWORD","root")
	cfg.Addr=config.GetString("DB_ADDR","127.0.0.1:3306")
	cfg.Net=config.GetString("DB_NET","tcp")
	cfg.DBName=config.GetString("DBName","auth_dev")

	fmt.Println("Connecting to Database:",cfg.DBName,cfg.FormatDSN())

	db,err:=sql.Open("mysql",cfg.FormatDSN())

	if err!=nil{
		fmt.Println("Error connecting to database",err)
		return nil,err
	}

	fmt.Println("Trying to connect to database...")
	pingErr:=db.Ping()
	if pingErr!=nil{
		fmt.Println("Error pinging database:",pingErr)
		return nil,pingErr
	}

	fmt.Println("Connected to database successfully:",cfg.DBName)

	return db,nil
}