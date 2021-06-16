package main

import (
	"DistributedComputingSystem/app/registry"
	"DistributedComputingSystem/app/encrypt"
	"DistributedComputingSystem/app/services"
	"context"
	"fmt"
	stlog "log"
)

func main() {
	host, port := "localhost", "6000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	var r registry.Registration
	r.ServiceName = registry.EncrptService
	r.ServiceURL = serviceAddress

	ctx, err := service.Start(context.Background(),
	r,
		host,
		port,
		encrypt.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}

	<-ctx.Done()
	fmt.Println("Shutting down encrypt service")
}