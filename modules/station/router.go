package station

import (
	"net/http"

	"github.com/al-chakim/jadwal-mrt/common/response"
	"github.com/gin-gonic/gin"
)

func Initiate(router *gin.RouterGroup){

	stationService := NewService()

	station := router.Group("/stations")

	station.GET("", func(c *gin.Context) {
		// code service
		GetAllStation(c, stationService)
	})
}

func GetAllStation(c *gin.Context, service Service) {
	datas, err := service.GetAllStation()

	if err != nil {
		// handle error
		c.JSON(http.StatusBadRequest, response.APIResponse{
			Success: false,
			Message: err.Error(),
			Data: nil,
		})
		return
	}

	// balikin response
	c.JSON(http.StatusOK, response.APIResponse{
		Success: true,
		Message: "Successfully get all station",
		Data: datas,
	})
}