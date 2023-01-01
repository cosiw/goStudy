package api

import (
	"2ND_PRACTICE/internal/global"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RequestURI struct {
	ImageId string `uri:"id"`
}

type RequestFile struct {
	File *multipart.FileHeader `form:"file"`
}

type RequestString struct {
	No    string `json: "no"`
	Image string `json: "image"`
}
type Response struct {
	Res string `json: "res"`
}

func CreateImage(c *gin.Context) {
	// image를 images폴더에 저장
	// req := RequestFile{}
	// if err := c.ShouldBind(&req); err != nil {
	// 	fmt.Println(err)
	// 	c.JSON(http.StatusBadRequest, Response{Res: "file is not valid"})

	// 	return
	// }

	// if err := c.SaveUploadedFile(req.File, "images/"+strconv.FormatInt(global.MaxImageNumber, 10)); err != nil {
	// 	fmt.Println(err)
	// 	c.JSON(http.StatusBadRequest, Response{Res: "file is not valid"})

	// 	return
	// }

	req := RequestString{}
	if err := c.ShouldBindUri(&req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res: "file is not valid"})
		return
	}
	fmt.Println("req.no", req.No)
	fmt.Println("req.image", req.Image)
	data := []byte(req.Image)

	if err := os.WriteFile("images/"+strconv.FormatInt(global.MaxImageNumber, 10), data, 0644); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res: "file is not valid"})
		return
	}

	global.MaxImageNumber++
	c.JSON(http.StatusOK, Response{Res: "success"})
}

func ReadImage(c *gin.Context) {
	reqURI := RequestURI{}
	if err := c.ShouldBindUri(&reqURI); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res: "request is not valid"})
		return
	}

	c.File("images/" + reqURI.ImageId) // status를 200으로 return함
}

func DeleteImage(c *gin.Context) {
	reqURI := RequestURI{}
	if err := c.ShouldBindUri(&reqURI); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res: "request is not valid"})
		return
	}
	err := os.Remove("images/" + reqURI.ImageId)
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, fmt.Sprintf("id is %s", reqURI.ImageId))

}

func UpdateImage(c *gin.Context) {
	reqURI := RequestURI{}
	req := RequestFile{}
	if err := c.ShouldBindUri(&reqURI); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res: "request is not valid"})
		return
	}
	if err := c.ShouldBind(&req); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res: "request is not valid"})
		return
	}

	if err := c.SaveUploadedFile(req.File, "images/"+reqURI.ImageId); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res: "file is not valid"})
		return
	}

	c.JSON(http.StatusOK, fmt.Sprintf("%s 파일 수정", reqURI.ImageId))

}
