package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"letters-api/internal/letter"
)

func Run() {
	c := letter.CreateController()
	srv := httpServer(c)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			return
		}
	}()
	log.Print("Server Started")

	<-done

	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := srv.Shutdown(ctx); err != nil {
		//log.Fatalf("Server Shutdown Failed:%+v", err)
		return
	}
	log.Print("Server Exited Properly")
}

func httpServer(c *letter.Controller) *http.Server {
	s := http.Server{Handler: router(c), Addr: ":3000"}
	return &s
}
