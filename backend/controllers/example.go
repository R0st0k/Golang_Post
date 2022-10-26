package controllers

import (
	"backend/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetExample godoc
// @Summary Insert and get example data from database
// @Produce json
// @Success 200 "OK"
// @Failure 500 "Fail"
// @Router  /example [get]
func GetExample(c *gin.Context) {
	s := new(models.Sending)
	err := s.InsertExample()
	if err != nil {
		c.Error(err)
		return
	}
	sendings, err := s.FindExample()
	if err != nil {
		c.Error(err)
		return
	}

	po := new(models.PostOffice)
	err = po.InsertExample()
	if err != nil {
		c.Error(err)
		return
	}
	postOffices, err := po.FindExample()
	if err != nil {
		c.Error(err)
		return
	}

	c.IndentedJSON(http.StatusOK, []interface{}{sendings, postOffices})
}
