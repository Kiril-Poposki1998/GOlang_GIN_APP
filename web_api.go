package main

import (
	"github.com/Kiril-Poposki1998/GOA/controller"
	"github.com/Kiril-Poposki1998/GOA/service"
	"github.com/gin-gonic/gin"
)

var (
	personService    service.Person_Service      = service.New()
	personController controller.PersonController = controller.New(personService)
)

func main() {
	server := gin.Default()
	server.GET("/person", func(ctx *gin.Context) {
		ctx.JSON(200, personController.FindAll())
	})

	server.POST("/person", func(ctx *gin.Context) {
		ctx.JSON(200, personController.Save(ctx))
	})
	server.Run("localhost:8080")
}
