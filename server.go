package main

import (
	"fmt"
	"github.com/bruteforce1414/simple-web-server-with-test/say"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)


func main() {

	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))


	var elementPath string;

	elementPath="/saygoodbye"
	http.HandleFunc(elementPath, say.Say)

	elementPath="/sayhello"
	http.HandleFunc(elementPath, say.Say)

	elementPath="/saynothing"
	http.HandleFunc(elementPath, say.Say)

	elementPath="/sayeverything"
	http.HandleFunc(elementPath, say.Say)

	fmt.Println("Server is listening...")

	e.GET("/dashboard", say.TasksGet)
/*	e.POST("/v1/tasks", taskHandler.TaskPost)
	e.PUT("/v1/tasks", taskHandler.TaskPut)
	e.GET("/v1/tasks/:task_id", taskHandler.TaskGet)
	e.DELETE("/v1/tasks/:task_id", taskHandler.TaskDelete)*/
	e.Logger.Fatal(e.Start(":8080"))}