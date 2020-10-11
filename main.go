package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Ping struct {
	Name string `json:"name"`
}

var messages = make(chan string)
var done = make(chan bool)

func main() {
	// defer worker.Wait()
	// for i := 0; i < 1000; i++ {
	// 	job := worker.Job{
	// 		Action: PrintPayload,
	// 		Payload: map[string]string{
	// 			"time": time.Now().String(),
	// 		},
	// 	}
	// 	job.Fire()
	// }

	go func() {
		for {
			select {
			case msg := <-messages:
				println(msg)
			case <-done:
				fmt.Println("done")
			}
		}
	}()

	e := router()
	e.Logger.Fatal(e.Start(":1323"))
}

func router() *echo.Echo {
	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		ping := new(Ping)
		if err := c.Bind(ping); err != nil {
			return c.String(http.StatusInternalServerError, "Hello, World!")
		}

		messages <- ping.Name
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/ping", func(c echo.Context) error {
		done <- true
		return c.String(http.StatusOK, "Hello, World!")
	})
	return e
}
