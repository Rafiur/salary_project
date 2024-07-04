package main

import (
	"fmt"
	"salary_project/db"
	"salary_project/handlers"

	"github.com/labstack/echo/v4"
)

func main()  {

	db, err := db.DBconnection()
	
	fmt.Println(err)

	e := echo.New()

	handlers.SetUpRouter(e, db)
	
	e.Logger.Fatal(e.Start(":4001"))
	
	
}