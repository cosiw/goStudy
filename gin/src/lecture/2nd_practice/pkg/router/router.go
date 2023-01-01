package router

import (
	"2ND_PRACTICE/pkg/api"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("internal/**/*")
	r.GET("/request", api.Index)
	r.POST("/images", api.CreateImage)
	r.GET("/images/:id", api.ReadImage)
	r.DELETE("/images/:id", api.DeleteImage)
	r.PUT("/images/:id", api.UpdateImage)
	//UPDATE, DELETE

	return r
}
