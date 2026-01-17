package main

import (
	// "log"
	// "net/http"

	"github.com/al-chakim/jadwal-mrt/modules/station"
	"github.com/gin-gonic/gin"
)

func main() {

	InitialRouter()

}

func InitialRouter(){
	router := gin.Default()
	api := router.Group("/v1/api")

	station.Initiate(api)

	router.Run(":8080")
}