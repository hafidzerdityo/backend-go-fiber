package api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"hafidzresttemplate.com/datastore"
	"hafidzresttemplate.com/services"
)

type ApiSetup struct {
    Logger *logrus.Logger
	Services *services.ServiceSetup
}

func NewApiSetup(db *gorm.DB)(apiSet ApiSetup) {
    loggerInit := logrus.New()
	apiSet = ApiSetup{
		Logger: loggerInit,
		Services: &services.ServiceSetup{
			Logger: loggerInit,
			Db: db,
			Datastore: &datastore.DatastoreSetup{
				Logger: loggerInit,
			},
		},
	}
    return 
}

func InitApi(db *gorm.DB)(router *gin.Engine) {
	router = gin.New()
 	router.Use(gin.Logger())
 	router.Use(gin.Recovery())
	apiSetup := NewApiSetup(db)

	apiSetup.Logger.Info("Setting up routes...")
	api := router.Group("/api/v1")
	{
		user := api.Group("/user_management")
		user.GET("/users", apiSetup.GetUsers)
		user.GET("/usersraw", apiSetup.GetUsersRaw)
		user.GET("/usersrawmap", apiSetup.GetUsersRawMap)
		user.POST("/user", apiSetup.CreateUser)
	}

	return 
}