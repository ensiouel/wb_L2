package main

import (
	"calendar/internal/config"
	"calendar/internal/service"
	"calendar/internal/storage"
	"calendar/internal/transport/http"
	"calendar/internal/transport/http/handler"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	httppkg "net/http"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	conf, err := config.New()
	if err != nil {
		log.Fatal(err)
		return
	}

	pool, err := pgxpool.New(ctx, fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", conf.Postgres.User, conf.Postgres.Password, conf.Postgres.Host, conf.Postgres.Port, conf.Postgres.DB))
	if err != nil {
		log.Fatal(err)
		return
	}

	err = pool.Ping(ctx)
	if err != nil {
		log.Println(err)
		return
	}

	eventStorage := storage.NewEventStorage(pool)
	eventService := service.NewEventService(eventStorage)
	eventHandler := handler.NewEventHandler(eventService)

	server := http.NewServer().Handle(eventHandler)
	go func() {
		err = server.Listen(conf.Server.Addr)
		if err != nil && err != httppkg.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	<-ctx.Done()
}
