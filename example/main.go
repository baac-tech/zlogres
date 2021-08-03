package main

import (
	"fmt"
	"regexp"

	"github.com/buildingwatsize/zlogres"
	"github.com/gofiber/fiber/v2"
)

func main() {

	// Default
	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	app.Use(zlogres.New())

	// [Custom] with requestid middleware
	/*
		app := fiber.New(fiber.Config{DisableStartupMessage: true})
		app.Use(requestid.New())
		app.Use(zlogres.New())
	*/

	// [Custom] with requestid middleware but use "transaction-id" as context locals key
	/*
		app := fiber.New(fiber.Config{DisableStartupMessage: true})
		app.Use(requestid.New(requestid.Config{
			ContextKey: "transaction-id",
		}))
		app.Use(zlogres.New(zlogres.Config{
			RequestIDContextKey: "transaction-id",
		}))
	*/

	app.Get("/", HandlerDefault)       // GET http://localhost:8000/
	app.Get("/msg/*", HandlerMsgParam) // GET http://localhost:8000/msg/{MESSAGE}

	fmt.Println("Listening on http://localhost:8000")
	fmt.Println("Try to send a request :D")
	app.Listen(":8000")
}

func HandlerDefault(c *fiber.Ctx) error {
	beautyCallLog("HandlerDefault")
	return c.SendString("Watch your app logs!")
}

func HandlerMsgParam(c *fiber.Ctx) error {
	beautyCallLog("HandlerMsgParam")

	msg := c.Params("*")
	c.Locals("message", msg) // Set context "message"

	return c.SendString("Watch your app logs! and see the difference (Hint: `message` will show on your logs)")
}

func beautyCallLog(called string) {
	m := regexp.MustCompile(".")
	dashed := "------------" + m.ReplaceAllString(called, "-") + "----"
	fmt.Println(dashed)
	fmt.Printf("--- Called: %v ---\n", called)
	fmt.Println(dashed)
}
