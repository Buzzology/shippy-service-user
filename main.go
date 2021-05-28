package main

import (
	pb "github.com/Buzzology/shippy-service-user/proto/user"
	"github.com/micro/go-micro/v2"
	// _ "github.com/micro/go-plugins/registry/mdns"
	// _ "github.com/asim/nitro-plugins/registry/mdns"
	"log"
)

func main() {

	// Create a database connection and close it before exit
	db, err := CreateConnection()
	defer db.Close() // Defer will be called when function returns

	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
	}

	// TODO: Look into how this works
	// Migrates the user struct into database column. Will check for changes
	// and migrate them each time the service is started.
	db.AutoMigration(&pb.User{})

	service := micro.NewService(

		// This name must match the package name given in protobuf definition
		micro.Name("shippy.service.user"),
	)

	repo := &UserRepository{db}

	// Init will parse the command line flags.
	srv.Init()

	// Register handler
	pb.RegisterAuthHandler(srv.Server(), &service{repo, tokenService})

	// Run the server
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
