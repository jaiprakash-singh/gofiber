package main

import (
	"context"
	"fmt"
	"gofiber/internal/db"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//app := fiber.New()

	// Custom config
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Test App v1.0.1",
	})

	if !fiber.IsChild() {
		fmt.Println("I'm the parent process")
	} else {
		fmt.Println("I'm a child process")
	}

	//Simple route: Respond with "Hello, World!" on root path, "/"
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello World!")
	// })
	app.Get("/", func(c *fiber.Ctx) error {
		return fiber.NewError(782, "Custom error message")
	})

	//Parameters: GET http://localhost:3000/say/hello%20world
	app.Get("/say/:value", func(c *fiber.Ctx) error {
		return c.SendString("value: " + c.Params("value"))
		// => Get request with value: hello world
	})

	//Optional parameter: GET http://localhost:3000/greet/john
	app.Get("/greet/:name?", func(c *fiber.Ctx) error {
		if c.Params("name") != "" {
			return c.SendString("Hello " + c.Params("name"))
			// => Hello john
		}
		return c.SendString("Where is john?")
	})

	//Wildcards: GET http://localhost:3000/api/user/jai
	// app.Get("/api/*", func(c *fiber.Ctx) error {
	// 	return c.SendString("API path: " + c.Params("*"))
	// 	// => API path: user/jai
	// })

	// http://localhost:3000/static/readme.txt
	app.Static("/static", "./public")

	//Route Handlers
	app.Get("/api/getName	", func(c *fiber.Ctx) error {
		return c.SendString("The name is fiber!")

		// => API path: user/jai
	})

	app.Put("/api/putName	", func(c *fiber.Ctx) error {
		return c.SendString("The name is fiber!")
		// => API path: user/jai
	})

	//app.Get("/api/connect", func(c *fiber.Ctx) error {
	if !fiber.IsChild() {
		func() {
			mongo := db.Mongo{}
			s := ""
			s += mongo.Connect()
			defer mongo.Client.Disconnect(context.TODO())
			s += mongo.List()
			mongo.Init()
			mongo.Insert()
			mongo.Find()
			mongo.Update()
			fmt.Println(s)
		}()
	}

	app.Listen(":3000")
}
