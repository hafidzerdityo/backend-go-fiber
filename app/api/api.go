package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
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

func InitApi(db *gorm.DB)(app *fiber.App) {
	app = fiber.New()
	app.Use(logger.New())

	apiSetup := NewApiSetup(db)
	apiSetup.Logger.Info("Setting up routes...")

	api := app.Group("/api")
	v1 := api.Group("/v1")
	user := v1.Group("/user_management")
	user.Get("/users", apiSetup.GetUsers)
	user.Get("/usersraw", apiSetup.GetUsersRaw)
	user.Get("/usersrawmap", apiSetup.GetUsersRawMap)
	user.Post("/user", apiSetup.CreateUser)
	

	return 
}