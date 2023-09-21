package services

import (
	"github.com/sirupsen/logrus"
	"hafidzresttemplate.com/datastore"
)


type ServiceSetup struct{
	Logger *logrus.Logger
	Datastore *datastore.DatastoreSetup
}



