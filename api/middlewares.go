package api

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

func setupMiddlewares(app *fiber.App) {
	// recover from panics
	app.Use(recover.New())

	// basic rate limiting
	app.Use(limiter.New(limiter.Config{
		Max:        30,
		Expiration: 1 * time.Minute,
	}))

	// request id
	app.Use(requestid.New())

	// logger
	app.Use(logger.New(logger.Config{
		Format:   "[${time}] REQ-ID=${locals:requestid} ${status} - ${method} ${path}\n",
		TimeZone: "UTC",
	}))
}
