package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"letters-api/internal/letter"
)

func Run() {

	stop := make(chan os.Signal, 1)
	kill := make(chan struct{}, 1)

	c := letter.CreateController(kill)
	srv := httpServer(c)

	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			fmt.Println(err)
			return
		}
	}()
	log.Print("Server Started")

	select {
	case <-stop:
		fmt.Println("os stop signal received")
	case <-kill:
		fmt.Println("kill signal received")
	}

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
	crs := routerWithCors(c)

	s := http.Server{Handler: crs, Addr: ":3000"}
	return &s
}
