package file

import (
	"errors"
	"net/http"
	"server/infrastructure/common/response"
	assembler "server/interfaces/assembler/file"
	dto "server/interfaces/dto/file"

	"github.com/gin-gonic/gin"
)

func (ctl *EndpointCtl) CreateFile(c *gin.Context) {
	var dtoFile dto.File
	err := c.ShouldBindJSON(&dtoFile)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Fail(errors.New("request body error")))
	}

	entityFile := assembler.DTOToEntity(&dtoFile)
	_, err = ctl.Srv.CreateFile(entityFile)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Fail(err))
		return
	}

	c.JSON(http.StatusCreated, response.Ok())
}
