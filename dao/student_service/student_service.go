package student_service

import (
	"rudn/dao/student_dao"
	"rudn/dao/user_interface"
	"strconv"
)

type StudentDAO interface {
	GetAllStudents() ([]student_dao.Student, error)
	FindStudentByID(id string) (student_dao.Student, error)
	AddStudent(student student_dao.Student) error
}

type StudentService struct {
	dao StudentDAO
}

func New(dao StudentDAO) *StudentService {
	return &StudentService{dao: dao}
}

func (s *StudentService) GetAllStudents() ([]user_interface.Student, error) {
	students, err := s.dao.GetAllStudents()

	var studentsUI []user_interface.Student
	for _, student := range students {
		studentUI, err := s.convertToUI(&student)
		if err != nil {
			return nil, err
		}
		studentsUI = append(studentsUI, *studentUI)
	}

	return studentsUI, err
}

func (s *StudentService) FindStudentByID(id int) (user_interface.Student, error) {
	student, err := s.dao.FindStudentByID(strconv.Itoa(id))
	if err != nil {
		return user_interface.Student{}, err
	}

	studentUI, err := s.convertToUI(&student)
	if err != nil {
		return user_interface.Student{}, err
	}

	return *studentUI, err
}

func (s *StudentService) AddStudent(student user_interface.Student) error {
	studentDAO := s.convertToDAO(&student)

	err := s.dao.AddStudent(*studentDAO)

	return err
}

func (s *StudentService) convertToUI(student *student_dao.Student) (*user_interface.Student, error) {
	id, err := strconv.Atoi(student.Id)
	if err != nil {
		return nil, err
	}
	points, err := strconv.Atoi(student.Points)
	if err != nil {
		return nil, err
	}

	studentUI := user_interface.Student{
		Id:      id,
		Name:    student.Name,
		Surname: student.Surname,
		Points:  points,
	}

	return &studentUI, nil
}

func (s *StudentService) convertToDAO(student *user_interface.Student) *student_dao.Student {
	studentDAO := student_dao.Student{
		Id:      strconv.Itoa(student.Id),
		Name:    student.Name,
		Surname: student.Surname,
		Points:  strconv.Itoa(student.Points),
	}

	return &studentDAO
}
