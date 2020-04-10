package config

import (
    "database/sql" 
    _ "github.com/go-sql-driver/mysql" 
    "os"  
)

var username string
var password string
var database string  
var url string
var port string 

func Connect() *sql.DB{

    username = os.Getenv("DB_USERNAME")
    password = os.Getenv("DB_PASSWORD")
    database = os.Getenv("DB_DATABASE")
    url = os.Getenv("DB_URL")
    port= os.Getenv("DB_PORT")

    conection := username+":"+password+"@tcp("+url+":"+port+")/"+database 

	DB, err := sql.Open("mysql", conection)
	if err != nil { 
        panic("failed to connect database "+err.Error())
    } 
    DB.SetMaxIdleConns(3)
    return DB
}