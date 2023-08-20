package main

import (
	"Test_Task_0/internal/nats"
	"Test_Task_0/internal/storage"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

import (
	"Test_Task_0/config"
	"Test_Task_0/internal/handler"
	"Test_Task_0/internal/repository"
	"Test_Task_0/internal/service"
	"context"
)

func main() {
	//natsClusterID := os.Getenv("NATS_CLUSTERID")
	//natsClientID := os.Getenv("NATS_CLIENTID")
	//natsURL := os.Getenv("NATS_URL")
	//natsChannel := os.Getenv("NATS_CHANNEL")
	//httpServerAddr := os.Getenv("HTTP_SERVER_ADDR")
	//host := os.Getenv("DB_HOST")
	//port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	//user := os.Getenv("DB_USER")
	//password := os.Getenv("DB_PASSWORD")
	//dbname := os.Getenv("DB_NAME")
	cfg := config.LoadConfig()
	logrus.Println("Configuration parsed successfully.")

	db := storage.ConnectToPostgres(cfg)
	logrus.Println("Database connected successfully.")

	defer storage.CloseDBConnection(db)

	repo := repository.NewRepository(db)
	handlers := handler.NewHandler(repo)

	sc := nats.NewNatsStream(cfg)
	sc.RunNatsSteaming(repo)

	server := new(service.Server)
	go func() {
		if err := server.Run(cfg, handlers.InitRoutes()); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("Error running http server: %s.", err.Error())
		}
	}()
	logrus.Print("Server started.")

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := server.ShutDown(context.Background()); err != nil {
		logrus.Errorf("Error shutting down: %s.", err.Error())
	}
	if err := sc.ShutDown(); err != nil {
		logrus.Errorf("Error nats streaming shutting down: %s.", err.Error())
	}
	//if err := db.Close(); err != nil {
	//	logrus.Errorf("Error db connection close : %s.", err.Error())
	//}
}
