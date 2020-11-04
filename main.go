package main

import (
	//"fmt"
	//"strings"
	"myapp/application"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:9000
// @BasePath /
func main() {
	//a := "2~23~45"
	//b := strings.Replace(a, "~", ",", -1)
	//fmt.Println("aaaa", a, b)
	application.Run()
}
