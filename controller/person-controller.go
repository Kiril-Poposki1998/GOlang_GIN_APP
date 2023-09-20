package controller

import (
	"net/http"

	"github.com/Kiril-Poposki1998/GOlang_GIN_APP/entity"
	"github.com/Kiril-Poposki1998/GOlang_GIN_APP/service"
	"github.com/gin-gonic/gin"
)

type PersonController interface {
	FindAll() []entity.Person
	Save(ctx *gin.Context) entity.Person
	ShowAll(ctx *gin.Context)
}

type controller struct {
	service service.Person_Service
}

func New(service service.Person_Service) PersonController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []entity.Person {
	return c.service.FindAll()
}

func (c *controller) Save(ctx *gin.Context) entity.Person {
	var person entity.Person
	ctx.BindJSON(&person)
	c.service.Save(person)
	return person
}

func (c *controller) ShowAll(ctx *gin.Context) {
	persons := c.service.FindAll()
	data := gin.H{
		"title":   "Person Page",
		"persons": persons,
	}
	ctx.HTML(http.StatusOK, "index.html", data)
}
