package api

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *ApiSetup) GetBooks(c *gin.Context) {
	a.Logger.Info("Getting books...")

	c.IndentedJSON(http.StatusOK, a.Services.GetBooks())
}

func (a *ApiSetup) GetTestMap(c *gin.Context) {
	a.Logger.Info("Getting test map...")

	jsonData, err := json.Marshal(a.Services.GetMap())
	if err != nil {
		a.Logger.Error(err)
		c.Error(err)
		return
	}

	c.Data(http.StatusOK, "application/json", jsonData)
}