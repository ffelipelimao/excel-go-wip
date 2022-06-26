package api

import (
	"net/http"

	"github.com/ffelipelimao/excel-go/internal/business"
	"github.com/gin-gonic/gin"
)

type handler struct {
	service business.GameService
}

func Setup(g *gin.Engine, s business.GameService) {

	handler := handler{
		service: s,
	}

	router := g.Group("/api/v1")

	router.GET("/games", handler.getGames)
	router.GET("/games/excel", handler.getGamesExcel)
	router.POST("/games/excel", handler.postGamesExcel)

}

func (h handler) getGames(c *gin.Context) {
	games, err := h.service.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, games)
}

func (h handler) getGamesExcel(c *gin.Context) {
	file, err := h.service.ListExcel()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.Header("Content-Disposition", "attachment; filename=data-20060102150405.xlsx")
	c.Data(http.StatusOK, "application/octet-stream", file.Bytes())
}

func (h handler) postGamesExcel(c *gin.Context) {
	formFile, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	fileOpen, _ := formFile.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	games, err := h.service.ReadExcel(fileOpen)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, games)
}
