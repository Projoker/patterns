package user_interface

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type mode int

const (
	MODE_SELECT mode = iota
	MODE_ADD
	MODE_FIND
)

type action int

const (
	ACTION_LIST action = iota
	ACTION_ADD
	ACTION_FIND
)

type StudentService interface {
	GetAllStudents() ([]Student, error)
	FindStudentByID(id int) (Student, error)
	AddStudent(student Student) error
}

type UserInterface struct {
	studentService StudentService
	mode           mode
}

type Student struct {
	Id      int
	Name    string
	Surname string
	Points  int
}

func New(service StudentService) *UserInterface {
	return &UserInterface{studentService: service, mode: MODE_SELECT}
}

func (ui *UserInterface) Run() {
	ui.printActions()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()

		text := scanner.Text()
		ui.handle(text)
	}
}

func (ui *UserInterface) handle(text string) {
	switch ui.mode {
	case MODE_SELECT:
		n, err := strconv.Atoi(text)
		if err != nil {
			ui.printActions()
			break
		}

		switch action(n) {
		case ACTION_LIST:
			students, err := ui.studentService.GetAllStudents()
			if err != nil {
				fmt.Println("get all students error", err)
				break
			}

			ui.printStudents(students)
			ui.back()
		case ACTION_ADD:
			ui.mode = MODE_ADD
			ui.printAdd()
		case ACTION_FIND:
			ui.mode = MODE_FIND
			ui.printFind()
		default:
			ui.printActions()
		}

	case MODE_ADD:
		fields := strings.Fields(text)

		if len(fields) != 4 {
			fmt.Print("parse student error, try again: ")
			break
		}

		id, err := strconv.Atoi(fields[0])
		if err != nil {
			fmt.Print("parse student error, try again: ")
			break
		}

		points, err := strconv.Atoi(fields[3])
		if err != nil {
			fmt.Print("parse student error, try again: ")
			break
		}

		student := Student{
			Id:      id,
			Name:    fields[1],
			Surname: fields[2],
			Points:  points,
		}
		err = ui.studentService.AddStudent(student)
		if err != nil {
			fmt.Println("add student error", err)
		}
		ui.back()

	case MODE_FIND:
		id, err := strconv.Atoi(text)
		if err != nil {
			fmt.Print("parse id error, try again: ")
			break
		}

		student, err := ui.studentService.FindStudentByID(id)
		if err != nil {
			fmt.Println("find student error: ", err)
		} else {
			ui.printStudent(&student)
		}

		ui.back()
	}
}

func (ui *UserInterface) back() {
	ui.printActions()
	ui.mode = MODE_SELECT
}

func (ui *UserInterface) printActions() {
	fmt.Println("Actions: ACTION_LIST - 0, ACTION_ADD - 1, ACTION_FIND - 2")
	fmt.Print("Choose action: ")
}

func (ui *UserInterface) printAdd() {
	fmt.Print("Enter student id name surname points by space: ")
}

func (ui *UserInterface) printFind() {
	fmt.Print("Enter student id: ")
}

func (ui *UserInterface) printStudents(students []Student) {
	fmt.Println("Total students:", len(students))

	for _, student := range students {
		ui.printStudent(&student)
	}
}

func (ui *UserInterface) printStudent(student *Student) {
	fmt.Printf("Id: %v, Name: %v, Surname: %v, Points: %v\n", student.Id, student.Name, student.Surname, student.Points)
}
