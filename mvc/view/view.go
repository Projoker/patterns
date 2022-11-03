package view

import (
	"bufio"
	"fmt"
	"os"
	"rudn/mvc/model"
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

type Controller interface {
	AddStudent(student model.Student) error
}

type Model interface {
	GetAllStudents() ([]model.Student, error)
	FindStudentByID(id string) (model.Student, error)
}

type View struct {
	m    Model
	c    Controller
	mode mode
}

func New(m Model, c Controller) *View {
	return &View{m: m, c: c, mode: MODE_SELECT}
}

func (ui *View) Run() {
	ui.printActions()

	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()

		text := scanner.Text()
		ui.handle(text)
	}
}

func (ui *View) handle(text string) {
	switch ui.mode {
	case MODE_SELECT:
		n, err := strconv.Atoi(text)
		if err != nil {
			ui.printActions()
			break
		}

		switch action(n) {
		case ACTION_LIST:
			students, err := ui.m.GetAllStudents()
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

		student := model.Student{
			Id:      fields[0],
			Name:    fields[1],
			Surname: fields[2],
			Points:  fields[3],
		}
		err := ui.c.AddStudent(student)
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

		student, err := ui.m.FindStudentByID(strconv.Itoa(id))
		if err != nil {
			fmt.Println("find student error: ", err)
		} else {
			ui.printStudent(&student)
		}

		ui.back()
	}
}

func (ui *View) back() {
	ui.printActions()
	ui.mode = MODE_SELECT
}

func (ui *View) printActions() {
	fmt.Println("Actions: ACTION_LIST - 0, ACTION_ADD - 1, ACTION_FIND - 2")
	fmt.Print("Choose action: ")
}

func (ui *View) printAdd() {
	fmt.Print("Enter student id name surname points by space: ")
}

func (ui *View) printFind() {
	fmt.Print("Enter student id: ")
}

func (ui *View) printStudents(students []model.Student) {
	fmt.Println("Total students:", len(students))

	for _, student := range students {
		ui.printStudent(&student)
	}
}

func (ui *View) printStudent(student *model.Student) {
	fmt.Printf("Id: %v, Name: %v, Surname: %v, Points: %v\n", student.Id, student.Name, student.Surname, student.Points)
}
