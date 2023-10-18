package datastore

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"hafidzresttemplate.com/dao"
	"hafidzresttemplate.com/data"
)

// func(d *DatastoreSetup) GetUsers(tx *gorm.DB)(datastoreResponse []data.GetUserQuerySelectItems, err error){

// 	sqlQuery := `SELECT * FROM public."user"`

//     // Execute the raw SQL query
//     rawQuery := tx.Raw(sqlQuery)

//     // Handle any query execution error
//     if rawQuery.Error != nil {
//         return nil, rawQuery.Error
//     }

//     // Scan the result into the datastoreResponse slice
//     if err := rawQuery.Scan(&datastoreResponse).Error; err != nil {
//         return datastoreResponse, err
//     }

// 	return
// }

func(d *DatastoreSetup) GetUsers(tx *gorm.DB)(datastoreResponse []dao.User, err error){
	d.Logger.Info(
		logrus.Fields{}, nil, "START: GetUsers Datastore",
	)

	selectedColumns := []string{
		"Username", "Nama", "Role", "Divisi", "Jabatan", "CreatedAt", "UpdatedAt", "IsDeleted",
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

func(d *DatastoreSetup) InsertUser(tx *gorm.DB, reqPayload dao.User)(datastoreResponse data.CreateUserResItems, err error){
	d.Logger.Info(
		logrus.Fields{}, nil, "START: InsertUser Datastore",
	)
	var sampleTrans dao.Transaksi

	sampleTrans.Username = reqPayload.Username

	if err := tx.Create(&sampleTrans).Error; err != nil {
		d.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, err.Error(),
		)
		return datastoreResponse, err
	}

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

