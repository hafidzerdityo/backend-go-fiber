package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"hafidzresttemplate.com/api"
)

func main() {
	godotenv.Load("service.env")
	host := os.Getenv("SERVICE_HOST")
	port := os.Getenv("SERVICE_PORT")
	
	app := api.InitApi()
	app.Run(fmt.Sprintf("%v:%v",host, port))
}
