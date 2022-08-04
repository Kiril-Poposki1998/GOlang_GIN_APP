package service

import "github.com/Kiril-Poposki1998/GOA/entity"

type Person_Service interface {
	Save(entity.Person) entity.Person
	FindAll() []entity.Person
}

type personService struct {
	persons []entity.Person
}

func New() Person_Service {
	return &personService{
		persons: []entity.Person{},
	}
}

func (service *personService) Save(person entity.Person) entity.Person {
	service.persons = append(service.persons, person)
	return person
}

func (service *personService) FindAll() []entity.Person {
	return service.persons
}
