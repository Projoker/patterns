package main

import (
	"rudn/dao/student_dao"
	"rudn/dao/student_service"
	"rudn/dao/user_interface"
	"rudn/mvc/controller"
	"rudn/mvc/model"
	"rudn/mvc/view"
)

func main() {
	//ctx := context.Background()
	//
	//m := model.New(ctx, "students.csv")
	//vm := viewmodel.New(ctx, m)
	//_ = view.New(ctx, vm)
	//
	//for {
	//	select {
	//	case <-ctx.Done():
	//		return
	//	default:
	//		time.Sleep(time.Second)
	//	}
	//}

	runDAO()
	//runMVC()
}

func runDAO() {
	dao := student_dao.New("students.csv")
	service := student_service.New(dao)
	ui := user_interface.New(service)

	ui.Run()
}

func runMVC() {
	m := model.New("students.csv")
	c := controller.New(m)
	v := view.New(m, c)

	v.Run()
}
