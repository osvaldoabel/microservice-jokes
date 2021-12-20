package main

import (
	"context"
	"fmt"
	stLog "log"

	"github.com/osvaldoabel/ms_app/grades"
	"github.com/osvaldoabel/ms_app/registry"
	"github.com/osvaldoabel/ms_app/service"
)

func main() {
	host, port := "localhost", "6000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	r := registry.Registration{
		ServiceName: registry.GradingService,
		ServiceURL:  serviceAddress,
	}

	ctx, err := service.Start(context.Background(),
		host,
		port,
		r,
		grades.RegisterHandlers,
	)

	if err != nil {
		stLog.Fatal(err)
	}
	<-ctx.Done()
	fmt.Println("shutting down grading service")

}
