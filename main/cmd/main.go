package main

import (
	"github.com/ffelipelimao/excel-go/internal/api"
	"github.com/ffelipelimao/excel-go/internal/business"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	service := business.NewGameServiceImpl()

	api.Setup(app, service)

	app.Run(":8080")
}
