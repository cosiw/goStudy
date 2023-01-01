package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name    string `form:"name" json: "name"`
	Address string `form:"address" json: "address"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	//router.StaticFile("/favicon.ico", "./resources/favicon.ico")
	//이미지만 불러온다(요청받는path, 이미지 위치 ) StaticFile은 Get 요청
	r.StaticFile("/image", "./cropped-webiste-banner.jpg")

	r.POST("/upload", func(c *gin.Context) {
		// Single file
		file, _ := c.FormFile("file")
		fmt.Println(file.Filename)

		// Upload the file to specific dst.
		// 저장하는 위치를 선택하려면 file.Filename 앞에 추가
		// c.SaveUploadedFile(file, "c://folder/ ~" + file.Filename)
		c.SaveUploadedFile(file, file.Filename)

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	r.GET("/test", func(c *gin.Context) {
		person := Person{}
		err := c.ShouldBind(&person)
		if err != nil {
			fmt.Println("Bind ERROR!")
		}
		fmt.Println(person)
		c.JSON(200, person)
	})
	r.Run(":8888") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
