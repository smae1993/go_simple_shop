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
	// ۱. بارگذاری تنظیمات
	cfg := config.Load()

	// ۲. اتصال به دیتابیس و اجرای مایگریشن‌ها
	db, err := database.NewSQLiteConnection(cfg.DBPath, cfg.MigrationPath)
	if err != nil {
		log.Fatalf("❌ error in db initialization: %v", err)
	}
	defer db.Close()

	// ۳. ساخت سرور و سیم‌کشی لایه‌ها
	srv := server.New(cfg, db)
	httpServer := srv.HTTPServer()

	// ۴. اجرای سرور در یک Goroutine برای پشتیبانی از Graceful Shutdown
	go func() {
		log.Printf("🚀 وبلاگ روی http://localhost%s روشن شد", cfg.Port)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("❌ error in running server: %v", err)
		}
	}()

	// ۵. منتظر ماندن برای سیگنال خاموش شدن (Ctrl+C)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("⏳ در حال خاموش کردن امن سرور...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	httpServer.Shutdown(ctx)
}
