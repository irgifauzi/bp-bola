package main

import (
	"log"

	"github.com/irgifauzi/bp-bola/config"
	
	"github.com/aiteung/musik"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/irgifauzi/bp-bola/url"


	"github.com/gofiber/fiber/v2"
   _"github.com/irgifauzi/bp-bola/docs"
)

// @title TES SWAGGER ULBI
// @version 1.0
// @description This is a sample swagger for Fiber

// @contact.name API Support
// @contact.url https://github.com/irgifauzi
// @contact.email 714220035@std.ulbi.ac.id

// @host uas-bola-919c3c7dac41.herokuapp.com
// @BasePath /
// @schemes https http

func main() {
	site := fiber.New(config.Iteung)
	site.Use(cors.New(config.Cors))
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}