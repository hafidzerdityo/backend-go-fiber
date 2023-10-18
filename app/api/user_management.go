package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"hafidzresttemplate.com/data"
)

func (a *ApiSetup) GetUsers(c *gin.Context) {
	a.Logger.Info(
		logrus.Fields{}, nil, "START: GetUsers API",
	)

	data, err := a.Services.GetUsers()
	if err != nil{
		a.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, err.Error(),
		)
		c.IndentedJSON(http.StatusNotFound, data)
		return
	}
	c.IndentedJSON(http.StatusOK, data)

	a.Logger.Info(
		logrus.Fields{}, nil, "END: GetUsers API",
	)

	return
}

func (a *ApiSetup) GetUsersRaw(c *gin.Context) {
	a.Logger.Info(
		logrus.Fields{}, nil, "START: GetUsersRaw API",
	)

	data, err := a.Services.GetUsersRaw()
	if err != nil{
		a.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, err.Error(),
		)
		c.IndentedJSON(http.StatusNotFound, data)
		return
	}
	c.IndentedJSON(http.StatusOK, data)

	a.Logger.Info(
		logrus.Fields{}, nil, "END: GetUsersRaw API",
	)

	return
}

func (a *ApiSetup) GetUsersRawMap(c *gin.Context) {
	a.Logger.Info(
		logrus.Fields{}, nil, "START: GetUsersRawMap API",
	)

	data, err := a.Services.GetUsersRawMap()
	if err != nil{
		a.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, err.Error(),
		)
		c.IndentedJSON(http.StatusNotFound, data)
		return
	}

    c.Header("Content-Type", "application/json")
    c.JSON(http.StatusOK, data)


	// jsonData, err := json.Marshal(data)
	// if err != nil {
	// 	a.Logger.Error(err)
	// 	c.Error(err)
	// 	return
	// }

	// c.Data(http.StatusOK, "application/json", jsonData)

	a.Logger.Info(
		logrus.Fields{}, nil, "END: GetUsersRawMap API",
	)

	return
}

func (a *ApiSetup) CreateUser(c *gin.Context) {
	a.Logger.Info(
		logrus.Fields{}, nil, "START: CreateUser API",
	)

	var reqPayload data.CreateUserReq

	if err := c.BindJSON(&reqPayload); err != nil {
		a.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, err.Error(),
		)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := a.Services.CreateUser(reqPayload)
	if err != nil{
		a.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, err.Error(),
		)
		c.IndentedJSON(http.StatusInternalServerError, data)
		return
	}
	c.IndentedJSON(http.StatusCreated, data)

	
	a.Logger.Info(
		logrus.Fields{}, nil, "END: CreateUser API",
	)
	return
}