package main

import (
	"dgiiScraper/scraper"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/:fiscalIdentity", func(ctx *fiber.Ctx) error {
		rnc, err := scraper.GetCompanyDataByRNC(ctx.Params("fiscalIdentity"))
		if err != nil {
			return err
		}

		return ctx.SendString(rnc)
	})

	err := app.Listen(":3000")
	if err != nil {
		return
	}
}
