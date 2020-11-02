package main

import (
	"context"
	"fmt"
	"github.com/ilovesusu/su-gin/config"
	"github.com/ilovesusu/su-gin/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	logo := `
		 ____         ____ _
		/ ___| _   _ / ___(_)_ __
		\___ \| | | | |  _| | '_ \
		 ___) | |_| | |_| | | | | |
		|____/ \__,_|\____|_|_| |_|
	
	`
	fmt.Println(logo)
	engine := router.InitRouter()

	srv := &http.Server{
		Addr:    config.SuServer.HttpIP + ":" + config.SuServer.HttpPort,
		Handler: engine,
	}
	go func() {
		// service connections
		log.Println(srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

}
