package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"hafidzresttemplate.com/api"
	"hafidzresttemplate.com/dao"
)


func main() {
	godotenv.Load("service.env")
	host := os.Getenv("SERVICE_HOST")
	port := os.Getenv("SERVICE_PORT")

	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")

	dbInit, err := dao.InitializeDB(
		dbUser,
		dbName,
		dbPassword,
		dbHost,
	)
	if err != nil{
		fmt.Println("Error When Initializing DB")
		return
	}

	
	app := api.InitApi(dbInit)
	app.Run(fmt.Sprintf("%v:%v",host, port))
}
