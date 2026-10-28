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

	queries := database.New(db)
	postService := service.NewPostService(queries)

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

func (s *Server) HTTPServer() *http.Server {
	return &http.Server{
		Addr:         s.cfg.Port,
		Handler:      s.registerRoutes(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}
}
