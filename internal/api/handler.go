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

}

func (h handler) getGames(c *gin.Context) {
	//TODO: Create errors massage
	games, err := h.service.List()

	if err != nil {
		return
	}

	c.JSON(http.StatusOK, games)
}

func (h handler) getGamesExcel(c *gin.Context) {
	//TODO: Create errors massage
	file, err := h.service.ListExcel()

	if err != nil {
		return
	}

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename=data-20060102150405.xlsx")
	c.Data(http.StatusOK, "application/octet-stream", file.Bytes())

}
