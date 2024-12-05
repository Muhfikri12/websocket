package main

import (
	"flag"
	"log"
	"project/infra"
	"project/routes"
)

func main() {
	migrateDb := flag.Bool("m", false, "use this flag to migrate database")
	seedDb := flag.Bool("s", false, "use this flag to seed database")
	flag.Parse()

	ctx, err := infra.NewServiceContext(*migrateDb, *seedDb)
	if err != nil {
		log.Fatal("can't init service context %w", err)
	}

	r := routes.NewRoutes(*ctx)

	if err = r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}

}
