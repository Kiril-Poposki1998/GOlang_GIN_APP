package controller

import (
	"github.com/Kiril-Poposki1998/smidGIN/entity"
	"github.com/Kiril-Poposki1998/smidGIN/service"
	"github.com/gin-gonic/gin"
)

type PersonController interface {
	FindAll() []entity.Person
	Save(ctx *gin.Context) entity.Person
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
