package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ivan-prykhodko/envoy-workshop/internal/product"
)

func main() {
	srv := product.NewHttpServer()

	go func() {
		if err := srv.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			srv.Logger.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown:", err)
	}

	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		log.Println("Shutdown timed out")
	}

	log.Println("Server exiting")
}
