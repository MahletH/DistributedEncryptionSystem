package main

import (
	"DistributedComputingSystem/app/registry"
	"DistributedComputingSystem/app/log"
	"DistributedComputingSystem/app/services"
	"context"
	"fmt"
	stlog "log"
)

func main() {
	log.Run("./app.log")

	host, port := "localhost", "5000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)
	var r registry.Registration
	r.ServiceName = registry.LogService
	r.ServiceURL = serviceAddress
	ctx, err := service.Start(context.Background(),
		r,
		host,
		port,
		log.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}

	<-ctx.Done()
	fmt.Println("Shutting down log service")
}