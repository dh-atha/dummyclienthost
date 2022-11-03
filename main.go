package main

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Webhook(c echo.Context) error {
	time.Sleep(3 * time.Second)
	return c.JSON(200, map[string]string{
		"status": "good",
	})
}

func WrongWebhook(c echo.Context) error {
	return c.JSON(500, map[string]string{
		"status": "internal server error",
	})
}

func BRIWebhook(c echo.Context) error {
	return c.JSON(200, map[string]string{
		"status": "good",
	})
}

func main()  {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.POST("/webhook", Webhook)
	e.POST("/webhook/wrong", WrongWebhook)
	e.POST(":service/messages/webhook/:msisdn", BRIWebhook)
	e.Logger.Fatal(e.Start(":3001"))
}