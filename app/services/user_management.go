package services

import "hafidzresttemplate.com/data"

func (s *ServiceSetup)GetUsers() (appResult []data.User, err error) {
	s.Logger.Info("Users Service exected")
	tx := s.Db.Begin()
	if tx.Error != nil {
		return appResult, tx.Error
	}
	
	appResult, err = s.Datastore.CheckUser(tx)
	if err != nil {
		tx.Rollback()
		s.Logger.Error("Check User Failed, Rollback.")
	}

	tx.Commit()
	return
}

func (s *ServiceSetup)GetMap()(appResult []map[string]interface{}, err error){

	s.Logger.Info("GetMap Service exected")
	appResult, err = s.Datastore.CreateUser()
	return
}
