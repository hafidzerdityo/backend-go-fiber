package datastore

import (
	"gorm.io/gorm"
	"hafidzresttemplate.com/data"
)

func(d *DatastoreSetup) CheckUser(tx *gorm.DB)(queryResult []data.User, err error){
	

	sqlQuery := `SELECT * FROM public."user"`

    // Execute the raw SQL query
    rawQuery := tx.Raw(sqlQuery)
    
    // Handle any query execution error
    if rawQuery.Error != nil {
        return nil, rawQuery.Error
    }
	
    // Scan the result into the queryResult slice
    if err := rawQuery.Scan(&queryResult).Error; err != nil {
        return nil, err
    }

	return
}

func(d *DatastoreSetup) CreateUser()(queryResult []map[string]interface{}, err error){
	queryResult  = []map[string]interface{}{
		{"hello":1},
		{"hello":2},
		{"hello":23},
	}

	return
}