package controller

import (
	"rudn/mvc/model"
)

type Model interface {
	AddStudent(student model.Student) error
}

type Controller struct {
	m Model
}

func New(m Model) *Controller {
	return &Controller{m: m}
}

func (s *Controller) AddStudent(student model.Student) error {
	err := s.m.AddStudent(student)

	return err
}
