package main

import (
	"context"

	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/smae1993/go_simple_shop/internal/config"
	"github.com/smae1993/go_simple_shop/internal/database"
	"github.com/smae1993/go_simple_shop/internal/server"
)

func main() {

	cfg := config.Load()

	db, err := database.NewSQLiteConnection(cfg.DBPath, cfg.MigrationPath)
	if err != nil {
		log.Fatalf("❌ error in db initialization: %v", err)
	}
	defer db.Close()

	srv := server.New(cfg, db)
	httpServer := srv.HTTPServer()

	go func() {
		log.Printf("🚀 وبلاگ روی http://localhost%s روشن شد", cfg.Port)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ error in running server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("⏳ Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	httpServer.Shutdown(ctx)
}
