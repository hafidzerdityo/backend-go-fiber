package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"hafidzresttemplate.com/dao"
)

func (a *ApiSetup) GetUsers(c *fiber.Ctx) error  {
	a.Logger.Info(
		logrus.Fields{}, nil, "START: GetUsers API",
	)

	data, err := a.Services.GetUsers()
	if err != nil{
		a.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, err.Error(),
		)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"resp_msg" : err.Error(),
			"resp_data" : data,
		})
	}

	a.Logger.Info(
		logrus.Fields{}, nil, "END: GetUsers API",
	)

	return c.JSON(fiber.Map{
		"resp_msg" : "Get Users data Succeed",
		"resp_data" : data,
	})
}

func (a *ApiSetup) GetUsersRaw(c *fiber.Ctx) error {
	a.Logger.Info(
		logrus.Fields{}, nil, "START: GetUsersRaw API",
	)

	data, err := a.Services.GetUsersRaw()
	if err != nil{
		a.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, err.Error(),
		)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"resp_msg" : err.Error(),
			"resp_data" : data,
		})
	}

	a.Logger.Info(
		logrus.Fields{}, nil, "END: GetUsersRaw API",
	)

	return c.JSON(fiber.Map{
		"resp_msg" : "Get Users data RAW Succeed",
		"resp_data" : data,
	})
}

func (a *ApiSetup) GetUsersRawMap(c *fiber.Ctx) error {
	a.Logger.Info(
		logrus.Fields{}, nil, "START: GetUsersRawMap API",
	)

	data, err := a.Services.GetUsersRawMap()
	if err != nil{
		a.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, err.Error(),
		)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"resp_msg" : err.Error(),
			"resp_data" : data,
		})
	}

	a.Logger.Info(
		logrus.Fields{}, nil, "END: GetUsersRawMap API",
	)

	return c.JSON(fiber.Map{
		"resp_msg" : "Get Users data RAW MAP Succeed",
		"resp_data" : data,
	})
}

func (a *ApiSetup) CreateUser(c *fiber.Ctx) error {
	a.Logger.Info(
		logrus.Fields{}, nil, "START: CreateUser API",
	)

	var reqPayload dao.CreateUserReq

	if err := c.BodyParser(&reqPayload); err != nil {
		a.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, err.Error(),
		)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"resp_msg" : err.Error(),
			"resp_data" : dao.CreateUserQuery{
				Success: false,
			},
		})
	}

	data, err := a.Services.CreateUser(reqPayload)
	if err != nil{
		a.Logger.Error(
			logrus.Fields{"error": err.Error()}, nil, err.Error(),
		)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"resp_msg" : err.Error(),
			"resp_data" : data,
		})
	}

	
	a.Logger.Info(
		logrus.Fields{}, nil, "END: CreateUser API",
	)
	return c.JSON(fiber.Map{
		"resp_msg" : "Create User Data Succeed",
		"resp_data" : data,
	})
}