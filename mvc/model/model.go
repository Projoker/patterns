package model

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type Student struct {
	Id      string
	Name    string
	Surname string
	Points  string
}

type Model struct {
	path string
}

func New(path string) *Model {
	return &Model{path: path}
}

// GetAllStudents - Список всех студентов
func (s *Model) GetAllStudents() ([]Student, error) {
	students := make([]Student, 0)

	f, err := os.Open(s.path)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if len(record) != 4 {
			return nil, err
		}

		student := Student{
			Id:      record[0],
			Name:    record[1],
			Surname: record[2],
			Points:  record[3],
		}

		students = append(students, student)
	}

	return students, nil
}

// FindStudentByID - Поиск по id
func (s *Model) FindStudentByID(id string) (Student, error) {
	f, err := os.Open(s.path)
	defer f.Close()

	if err != nil {
		return Student{}, err
	}

	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return Student{}, err
		}

		if len(record) != 4 {
			return Student{}, err
		}

		if record[0] != id {
			continue
		}

		student := Student{
			Id:      record[0],
			Name:    record[1],
			Surname: record[2],
			Points:  record[3],
		}
		return student, nil
	}

	return Student{}, fmt.Errorf("not found")
}

// AddStudent Добавление в базу
func (s *Model) AddStudent(student Student) error {
	f, err := os.OpenFile(s.path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	defer f.Close()
	if err != nil {
		return err
	}

	data := []string{student.Id, student.Name, student.Surname, student.Points}

	w := csv.NewWriter(f)
	err = w.Write(data)
	if err != nil {
		return err
	}
	w.Flush()

	return nil
}
