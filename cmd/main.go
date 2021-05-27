package main

import (
	"github.com/LKarlon/converter/conv"
	"github.com/LKarlon/converter/pkg/handler"
	"github.com/LKarlon/converter/pkg/service"
	"github.com/sirupsen/logrus"
)

func main() {

	port := "80"
	service := service.NewService()
	handlers := handler.NewHandler(service)
	srv := new(conv.Server)
	go func() {
		if err := srv.Run(port, handlers.InitRoutes()); err != nil {
			logrus.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()
	logrus.Print("TodoApp Started")

	logrus.Print("TodoApp Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("TodoApp Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}
}
