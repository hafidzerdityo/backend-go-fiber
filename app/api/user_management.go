package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *ApiSetup) GetBooks(c *gin.Context) {
	a.Logger.Info("Getting books...")

	data, err := a.Services.GetBooks()
	if err != nil{
		
	}
	c.IndentedJSON(http.StatusOK, data)
}

func (a *ApiSetup) GetTestMap(c *gin.Context) {
	a.Logger.Info("Getting test map...")

	data, err := a.Services.GetMap()

	jsonData, err := json.Marshal(data)
	if err != nil {
		a.Logger.Error(err)
		c.Error(err)
		return
	}

	c.Data(http.StatusOK, "application/json", jsonData)
}