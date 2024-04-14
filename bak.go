package main

// import (
// 	"log"
// 	"net/http"
// 	"os"
// 	conn "testapi/connection"

// 	"github.com/labstack/echo/v4"
// 	"github.com/labstack/echo/v4/middleware"
// )

// func main() {

// 	e := echo.New()

// 	e.Use(middleware.Logger())
// 	e.Use(middleware.Recover())

// 	_, err := conn.New(&conn.Config{
// 		Host:         "localhost",
// 		Port:         10023,
// 		User:         "testUser",
// 		Password:     "P@ssw0rd",
// 		DatabaseName: "test",
// 		Debug:        true,
// 	})
// 	if err != nil {
// 		log.Println("Postgres was failed, %v", err.Error())
// 	}

// 	e.GET("/", func(c echo.Context) error {
// 		return c.HTML(http.StatusOK, "Hello, Docker! <3\n")
// 	})

// 	httpPort := os.Getenv("HTTP_PORT")
// 	if httpPort == "" {
// 		httpPort = "1323"
// 	}

// 	e.Logger.Fatal(e.Start(":" + httpPort))
// }
