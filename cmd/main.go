package main

import (
	"context"
	"fmt"
	"github.com/Sokol111/educational-go/internal/config"
	"github.com/Sokol111/educational-go/internal/handler"
	"github.com/Sokol111/educational-go/internal/repository"
	"github.com/Sokol111/educational-go/internal/repository/db"
	"github.com/Sokol111/educational-go/internal/server"
	"github.com/Sokol111/educational-go/internal/service"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	conf := config.LoadConfig("./configs")
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	conn, err := pgx.Connect(ctx, fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		conf.DB.Host, conf.DB.Port, conf.DB.Username, conf.DB.Password, conf.DB.DBName, conf.DB.SSLMode))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)
	queries := db.New(conn)
	ur := repository.NewUserPostgresRepository(queries)
	us := service.NewUserService(ur)
	uh := handler.NewUserHandler(us)
	server.NewServer(conf.Port, ctx, uh).Start()
}
