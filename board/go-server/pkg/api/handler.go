package api

import (
	"fmt"
	"go-server/pkg/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Res string `json:"res"`
}

type RequestURI struct {
	Id string `uri:"id"`
}

func (apis *APIs) CreateBoard(c *gin.Context) {
	req := &model.Board{}

	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, &Response{Res: "Bad Request"})

		return
	}

	res, err := apis.db.CreateBoard(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{Res: "Bad request"})
		return
	}

	c.JSON(http.StatusOK, res)
}

func (apis *APIs) GetBoardList(c *gin.Context) {
	res, err := apis.db.GetBoardList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{Res: "SErver error"})

		return
	}

	c.JSON(http.StatusOK, res)
}

func (apis *APIs) DeleteBoard(c *gin.Context) {
	id := RequestURI{}
	fmt.Println(id.Id)

	if err := c.ShouldBindUri(&id); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, Response{Res: "request is not valid"})
		return
	}
	err := apis.db.DeleteBoardByID(id.Id)

	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{Res: "Bad request"})
		return
	}

	c.JSON(http.StatusOK, &Response{Res: "Success"})
}

func (apis *APIs) UpdateBoard(c *gin.Context) {
	req := &model.Board{}

	id := req.ID

	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, &Response{Res: "Bad request"})
		return
	}

	res, err := apis.db.UpdateBoard(uint(id), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, &Response{Res: "Bad request"})
		return
	}

	c.JSON(http.StatusOK, res)
}
