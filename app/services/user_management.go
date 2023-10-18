package services

import (
	"time"

	"github.com/jinzhu/copier"
	"github.com/sirupsen/logrus"
	"hafidzresttemplate.com/dao"
	"hafidzresttemplate.com/data"
	"hafidzresttemplate.com/pkg/utils"
)

func (s *ServiceSetup)GetUsers() (appResponse data.GetUserRes, err error) {
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
		appResponse.RespMsg = err.Error()
		return
	}

	err = copier.Copy(&appResponse.RespData, datastoreResponse)
	if err != nil {
		tx.Rollback()
		s.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, err.Error(),
		)
		appResponse.RespMsg = err.Error()
		return
	}

	appResponse.RespMsg = "Get Users data Succeed"

	tx.Commit()
	s.Logger.Info(
		logrus.Fields{}, nil, "END: GetUsers Service",
	)
	return
}

func (s *ServiceSetup)GetUsersRaw() (appResponse data.GetUserRes, err error) {
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
		appResponse.RespMsg = err.Error()
		return
	}

	appResponse.RespData = datastoreResponse

	appResponse.RespMsg = "Get Users data Succeed"

	tx.Commit()
	s.Logger.Info(
		logrus.Fields{}, nil, "END: GetUsersRaw Service",
	)
	return

}
func (s *ServiceSetup)GetUsersRawMap() (appResponse data.GetUserResRawMap, err error) {
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
		appResponse.RespMsg = err.Error()
		return
	}

	appResponse.RespData = datastoreResponse

	appResponse.RespMsg = "Get Users data Succeed"

	tx.Commit()
	s.Logger.Info(
		logrus.Fields{}, nil, "END: GetUsersRawMap Service",
	)
	return
}

func (s *ServiceSetup)CreateUser(reqPayload data.CreateUserReq) (appResponse data.CreateUserRes, err error) {
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
		appResponse.RespMsg = err.Error()
		return
	}

	reqPayloadToInsert.HashedPassword = getHashedPass
	reqPayloadToInsert.Nama = reqPayload.Nama
	reqPayloadToInsert.Role = reqPayload.Role
	reqPayloadToInsert.Divisi = reqPayload.Divisi
	reqPayloadToInsert.Jabatan = reqPayload.Jabatan
	reqPayloadToInsert.CreatedAt = time.Now()
	reqPayloadToInsert.UpdatedAt = nil
	reqPayloadToInsert.IsDeleted = false

	
	datastoreResponse, err := s.Datastore.InsertUser(tx, reqPayloadToInsert)
	if err != nil {
		tx.Rollback()
		s.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, err.Error(),
		)
		appResponse.RespMsg = err.Error()
		return
	}

	appResponse.RespData = datastoreResponse
	appResponse.RespMsg = "Success"

	tx.Commit()

	s.Logger.Info(
		logrus.Fields{}, nil, "END: CreateUser Service",
	)
	return
}

