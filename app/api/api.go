package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"hafidzresttemplate.com/services"
)

type ApiSetup struct {
    Logger *logrus.Logger
	Services *services.ServiceSetup
}

func NewApiSetup()(loggerSet ApiSetup) {
    loggerInit := logrus.New()
	loggerSet = ApiSetup{
		Logger: loggerInit}
    return 
}

func SetupRoutes()(router *gin.Engine) {
	router = gin.New()
 	router.Use(gin.Logger())
 	router.Use(gin.Recovery())

	apiSetup := NewApiSetup()
	apiSetup.Logger.Info("Setting up routes...")

	api := router.Group("/api/v1")
	{
		user := api.Group("/user_management")
		user.GET("/books", apiSetup.GetBooks)
		user2 := api.Group("/user_management2")
		user2.GET("/random_map", apiSetup.GetTestMap)
	}

	return 
}