package main

import (
	"flag"
	"log"
	_ "project/docs"
	"project/infra"
	"project/routes"
)

// @title Ecommerce Dashboard API
// @version 1.0
// @description Nothing.
// @termsOfService http://example.com/terms/
// @contact.name Team-1
// @contact.url https://academy.lumoshive.com/contact-us
// @contact.email lumoshive.academy@gmail.com
// @license.name Lumoshive Academy
// @license.url https://academy.lumoshive.com
// @host localhost:8080
// @schemes http
// @BasePath /
// @securityDefinitions.apikey token
// @in header
// @name token
func main() {
	migrateDb := flag.Bool("m", false, "use this flag to migrate database")
	seedDb := flag.Bool("s", false, "use this flag to seed database")
	flag.Parse()

	ctx, err := infra.NewServiceContext(*migrateDb, *seedDb)
	if err != nil {
		log.Fatal("can't init service context %w", err)
	}

	r := routes.NewRoutes(*ctx)

	if err = r.Run(":8000"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

}
