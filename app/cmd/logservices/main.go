package main

import (
	"DistributedComputingSystem/app/log"
	"DistributedComputingSystem/app/services"
	"context"
	"fmt"
	stlog "log"
)

func main() {
	log.Run("./app.log")

	host, port := "localhost", "5000"

	ctx, err := service.Start(context.Background(),
		"Log Service",
		host,
		port,
		log.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}

	<-ctx.Done()
	fmt.Println("Shutting down log service")
}