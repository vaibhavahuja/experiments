package main

import (
	"github.com/experimentsGo/entities"
	"github.com/experimentsGo/service"
)

func main() {
	template := entities.GetNewTemplate("2179e172asd ", "Vaibhav ")
	//rule := entities.Rule{
	//	Uuid: "asas",
	//	Name: "asasa",
	//}
	application := service.GetNewApplication(template)

	application.ConvertToRule()
}
