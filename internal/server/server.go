package server

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/smae1993/go_simple_shop/internal/config"
	"github.com/smae1993/go_simple_shop/internal/database"
	"github.com/smae1993/go_simple_shop/internal/handler"
	"github.com/smae1993/go_simple_shop/internal/service"
)

type Server struct {
	cfg           *config.Config
	db            *sql.DB
	publicHandler *handler.PublicHandler
	adminHandler  *handler.AdminHandler
}

func New(cfg *config.Config, db *sql.DB) *Server {
	// ۱. ساخت لایه دیتابیس و سرویس‌ها
	queries := database.New(db)
	postService := service.NewPostService(queries)

	// ۲. ساخت موتور قالب و کنترلرها
	renderer := handler.NewRenderer()
	publicHandler := handler.NewPublicHandler(postService, renderer)
	adminHandler := handler.NewAdminHandler(postService, renderer)

	return &Server{
		cfg:           cfg,
		db:            db,
		publicHandler: publicHandler,
		adminHandler:  adminHandler,
	}
}

// ساخت http.Server استاندارد با تنظیمات امنیتی Timeout برای محیط پروداکشن
func (s *Server) HTTPServer() *http.Server {
	return &http.Server{
		Addr:         s.cfg.Port,
		Handler:      s.registerRoutes(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
}
