package file

import (
	"errors"
	"net/http"
	"server/infrastructure/common/response"
	dto "server/interfaces/dto/file"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (ctl *EndpointCtl) GetFileById(c *gin.Context) {
	userID, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Fail(errors.New("parameter error")))
		return
	}
	user, err := ctl.Srv.GetFileById(int64(userID))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, response.Fail(err))
		return
	}
	fileinfo := dto.FileInfo{
		Filename: user.Name,
		Filetype: user.Type,
	}
	c.JSON(http.StatusOK, response.Data(fileinfo))
}
