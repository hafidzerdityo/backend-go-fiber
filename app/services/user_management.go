package services

import (
	"fmt"
	"time"

	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	"hafidzresttemplate.com/dao"
	"hafidzresttemplate.com/pkg/utils"
)

func (s *ServiceSetup)GetUsers() (appResponse []dao.GetUserQuery, err error) {
	s.Logger.Info(
		logrus.Fields{}, nil, "START: GetUsers Service",
	)
	tx := s.Db.Begin()
	if tx.Error != nil {
		return appResponse, tx.Error
	}
	
	datastoreResponse, err := s.Datastore.GetUsers(tx)
	if err != nil {
		tx.Rollback()
		s.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, err.Error(),
		)
		return
	}

	if len(datastoreResponse) == 0{
		tx.Rollback()
		err = fmt.Errorf("No Data Found")
		s.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, err.Error(),
		)
		return
	}

	err = copier.Copy(&appResponse, datastoreResponse)
	if err != nil {
		tx.Rollback()
		s.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, err.Error(),
		)
		return
	}

	tx.Commit()
	s.Logger.Info(
		logrus.Fields{}, nil, "END: GetUsers Service",
	)
	return
}

func (s *ServiceSetup)GetUsersRaw() (appResponse []dao.GetUserQuery, err error) {
	s.Logger.Info(
		logrus.Fields{}, nil, "START: GetUsersRaw Service",
	)
	tx := s.Db.Begin()
	if tx.Error != nil {
		return appResponse, tx.Error
	}
	
	datastoreResponse, err := s.Datastore.GetUsersRaw(tx)
	if err != nil {
		tx.Rollback()
		s.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, err.Error(),
		)

		return
	}

	appResponse = datastoreResponse

	tx.Commit()
	s.Logger.Info(
		logrus.Fields{}, nil, "END: GetUsersRaw Service",
	)
	return

}
func (s *ServiceSetup)GetUsersRawMap() (appResponse []map[string]interface{}, err error) {
	s.Logger.Info(
		logrus.Fields{}, nil, "START: GetUsersRawMap Service",
	)
	tx := s.Db.Begin()
	if tx.Error != nil {
		return appResponse, tx.Error
	}
	
	datastoreResponse, err := s.Datastore.GetUsersRawMap(tx)
	if err != nil {
		tx.Rollback()
		s.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, err.Error(),
		)
		return
	}

	appResponse = datastoreResponse

	tx.Commit()
	s.Logger.Info(
		logrus.Fields{}, nil, "END: GetUsersRawMap Service",
	)
	return
}

func (s *ServiceSetup)CreateUser(reqPayload dao.CreateUserReq) (appResponse dao.CreateUserQuery, err error) {
	s.Logger.Info(
		logrus.Fields{}, nil, "START: CreateUser Service",
	)
	tx := s.Db.Begin()
	if tx.Error != nil {
		return appResponse, tx.Error
	}

	var reqPayloadToInsert dao.User

	reqPayloadToInsert.Username = reqPayload.Username

	getHashedPass, err := utils.HashPassword(reqPayload.Password)
	if err != nil {
		tx.Rollback()
		s.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, err.Error(),
		)
		return
	}

	reqPayloadToInsert.HashedPassword = getHashedPass
	reqPayloadToInsert.Email = reqPayload.Email
	reqPayloadToInsert.Nama = reqPayload.Nama
	reqPayloadToInsert.Role = reqPayload.Role
	reqPayloadToInsert.CreatedAt = time.Now()
	reqPayloadToInsert.UpdatedAt = nil
	reqPayloadToInsert.IsDeleted = false

	
	datastoreResponse, err := s.Datastore.InsertUser(tx, reqPayloadToInsert)
	if err != nil {
		tx.Rollback()
		s.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, err.Error(),
		)
		return
	}

	appResponse = datastoreResponse

	tx.Commit()

	s.Logger.Info(
		logrus.Fields{}, nil, "END: CreateUser Service",
	)
	return
}

