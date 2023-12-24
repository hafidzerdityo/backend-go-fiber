package datastore

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"hafidzresttemplate.com/dao"
	"hafidzresttemplate.com/data"
)

func(d *DatastoreSetup) GetUsersRaw(tx *gorm.DB)(datastoreResponse []data.GetUserQuery, err error){
	d.Logger.Info(
		logrus.Fields{}, nil, "START: GetUsersRaw Datastore",
	)

	sqlQuery := `SELECT * FROM public."user"`

    rawQuery := tx.Raw(sqlQuery)

    if rawQuery.Error != nil {
		d.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, err.Error(),
		)
        return nil, rawQuery.Error
    }

    // Scan the result into the datastoreResponse slice
    if err := rawQuery.Scan(&datastoreResponse).Error; err != nil {
		d.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, err.Error(),
		)
        return datastoreResponse, err
    }

	d.Logger.Info(
		logrus.Fields{}, nil, "END: GetUsersRaw Datastore",
	)

	return
}

func(d *DatastoreSetup) GetUsersRawMap(tx *gorm.DB)(datastoreResponse []map[string]interface{}, err error){
	d.Logger.Info(
		logrus.Fields{}, nil, "START: GetUsersRawMap Datastore",
	)

	sqlQuery := `SELECT * FROM public."user"`
    rawQuery := tx.Raw(sqlQuery)

    if rawQuery.Error != nil {
		d.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, err.Error(),
		)
        return nil, rawQuery.Error
    }

    // Find the result into the datastoreResponse slice
    if err = rawQuery.Find(&datastoreResponse).Error; err != nil {
        return
    }
	
	for _, val := range(datastoreResponse){
		delete(val, "hashed_password")
	}

	d.Logger.Info(
		logrus.Fields{}, nil, "END: GetUsersRawMap Datastore",
	)

	return
}

func(d *DatastoreSetup) GetUsers(tx *gorm.DB)(datastoreResponse []dao.User, err error){
	d.Logger.Info(
		logrus.Fields{}, nil, "START: GetUsers Datastore",
	)

	selectedColumns := []string{
		"Username", "Nama", "Email" ,"Role", "CreatedAt", "UpdatedAt", "IsDeleted",
	}
    if err := tx.Select(selectedColumns).Find(&datastoreResponse).Error; err != nil {
		d.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, err.Error(),
		)
        return nil, err
    }

	d.Logger.Info(
		logrus.Fields{}, nil, "END: GetUsers Datastore",
	)

    return
}

func(d *DatastoreSetup) InsertUser(tx *gorm.DB, reqPayload dao.User)(datastoreResponse data.CreateUserQuery, err error){
	d.Logger.Info(
		logrus.Fields{}, nil, "START: InsertUser Datastore",
	)

	if err := tx.Create(&reqPayload).Error; err != nil {
		d.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, err.Error(),
		)
		return datastoreResponse, err
	}

	datastoreResponse.Success = true

	d.Logger.Info(
		logrus.Fields{}, nil, "END: InsertUser Datastore",
	)

	return
}

