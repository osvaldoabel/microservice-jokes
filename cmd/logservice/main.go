package main

import (
	"context"
	"fmt"
	stLog "log"

	"github.com/osvaldoabel/ms_app/log"
	"github.com/osvaldoabel/ms_app/registry"
	"github.com/osvaldoabel/ms_app/service"
)

func main() {
	log.Run("./app.log")
	host, port := "localhost", "4000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	var r registry.Registration
	r.ServiceName = registry.LogService
	r.ServiceURL = serviceAddress

	ctx, err := service.Start(
		context.Background(),
		host,
		port,
		r,
		log.RegisterHandlers,
	)

	if err != nil {
		stLog.Fatal(err)
	}
	<-ctx.Done()

	fmt.Println("Shutting down log service")
}
