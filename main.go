package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	fmt.Println("[DAEMON] system time is" + time.Now().Format(time.RFC850))

	app := fiber.New()

	fmt.Println("[DAEMON] app started on port 3000")

	app.Static("/assets", "./assets")
	app.Static("/", "./public/index.html")
	app.Static("*", "./public/error.html")

	app.Listen(":3000")
}
