package router

import (
	"go-server/pkg/api"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router(apis *api.APIs) *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())
	r.POST("/crud", apis.CreateBoard)
	r.GET("/crud", apis.GetBoardList)
	r.DELETE("/crud", apis.DeleteBoard)
	r.PUT("/crud", apis.UpdateBoard)

	return r
}
