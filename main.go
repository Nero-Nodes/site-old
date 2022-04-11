// -------------------- //
// nero-nodes/site      //
// /main.go             //
// (c) 2022 Nero        //
// http://neronodes.net //
// -------------------- //

package main

// Import our libraries needed for this project.
import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Create a main function to start the webserver instance.
func main() {
	// Print system time to console for debug purposes.
	fmt.Println("[DAEMON] system time is " + time.Now().Format(time.RFC850))

	// Create the Fiber webserver instance.
	app := fiber.New()

	// Create the routes for pages and assets.
	app.Static("/assets", "./assets")
	app.Static("/", "./public/index.html")
	app.Static("/cancel", "./public/cancel.html")
	app.Static("*", "./public/error.html")

	// Listen on port 3000 for incoming requests.
	app.Listen(":3000")

	// Log this event to the console.
	fmt.Println("[DAEMON] app started on port 3000")
}
