package router

import "github.com/gofiber/fiber/v2"

func HealthRouter(router fiber.Router) {

	router.Get("/health", healthHandlers)
	router.Get("/status", sendStatusHandlers)
}

func healthHandlers(ctx *fiber.Ctx) error {

	_ = ctx.JSON(fiber.Map{"success": true})

	return nil
}

func sendStatusHandlers(ctx *fiber.Ctx) error {

	_ = ctx.SendStatus(400)

	return nil
}
