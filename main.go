package main

import (
	"log"

	"github.com/aiteung/musik"
	"github.com/xshazks/websocket-backend/module"
	"github.com/xshazks/websocket-backend/url"

	"github.com/gofiber/fiber/v2"
)

func main() {
	go module.RunHub()

	site := fiber.New()
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}