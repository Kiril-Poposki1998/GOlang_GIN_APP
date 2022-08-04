package main

import (
	"github.com/Kiril-Poposki1998/smidGIN/controller"
	"github.com/Kiril-Poposki1998/smidGIN/middleware"
	"github.com/Kiril-Poposki1998/smidGIN/service"
	"github.com/gin-gonic/gin"
)

var (
	personService    service.Person_Service      = service.New()
	personController controller.PersonController = controller.New(personService)
)

func main() {
	server := gin.New()
	server.Use(gin.Recovery(), middleware.BasicAuth())
	server.GET("/person", func(ctx *gin.Context) {
		ctx.JSON(200, personController.FindAll())
	})

	server.POST("/person", func(ctx *gin.Context) {
		ctx.JSON(200, personController.Save(ctx))
	})
	server.Run("localhost:8080")
}
