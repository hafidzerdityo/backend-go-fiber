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
	godotenv.Load("config.env")
	host := os.Getenv("SERVICE_HOST")
	port := os.Getenv("SERVICE_PORT")

	dbUser := os.Getenv("POSTGRES_USER")
	dbName := os.Getenv("POSTGRES_DB")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")

	dbInit, err := dao.InitializeDB(
		dbUser,
		dbName,
		dbPassword,
		dbHost,
		dbPort,
	)
	if err != nil{
		fmt.Println("Error When Initializing DB")
		return
	}

	
	app := api.InitApi(dbInit)
	app.Listen(fmt.Sprintf("%v:%v",host, port))
	if err != nil {
		fmt.Println("Error When Initializing Server")
		panic(err)
	}
}
