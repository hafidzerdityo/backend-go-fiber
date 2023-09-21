package services

import "hafidzresttemplate.com/data"

func (s *ServiceSetup)GetBooks() (appResult []data.Book, err error) {

	s.Logger.Info("Book Service exected")
	appResult, err = s.Datastore.CheckUser()
	return
}

func (s *ServiceSetup)GetMap()(appResult []map[string]interface{}, err error){

	s.Logger.Info("GetMap Service exected")
	appResult, err = s.Datastore.CreateUser()
	return
}
