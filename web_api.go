package main

import (
	"os"

	"github.com/Kiril-Poposki1998/smidGIN/controller"
	"github.com/Kiril-Poposki1998/smidGIN/service"
	"github.com/gin-gonic/gin"
)

var (
	personService    service.Person_Service      = service.New()
	personController controller.PersonController = controller.New(personService)
)

func main() {
	server := gin.New()
	server.LoadHTMLGlob("/app/templates/*.html")
	server.Use(gin.Recovery())

	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/person", func(ctx *gin.Context) {
			ctx.JSON(200, personController.FindAll())
		})

		apiRoutes.POST("/person", func(ctx *gin.Context) {
			ctx.JSON(200, personController.Save(ctx))
		})
	}

	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/persons", personController.ShowAll)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	server.Run(":" + port)
}
