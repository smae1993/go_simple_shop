package server

import (
	"net/http"

	"github.com/smae1993/go_simple_shop/internal/middleware"
)

func (s *Server) registerRoutes() http.Handler {
	mux := http.NewServeMux()

	// ۱. روت‌های عمومی سایت (Public)
	mux.HandleFunc("GET /", s.publicHandler.Home)
	mux.HandleFunc("GET /blog", s.publicHandler.BlogIndex)
	mux.HandleFunc("GET /blog/{slug}", s.publicHandler.SinglePost)

	// ۲. روت‌های پنل مدیریت (Admin)
	mux.HandleFunc("GET /admin/blog", s.adminHandler.List)
	mux.HandleFunc("GET /admin/blog/post", s.adminHandler.PostForm)
	mux.HandleFunc("POST /admin/blog/post", s.adminHandler.SavePost)
	mux.HandleFunc("POST /admin/blog/post/{id}/delete", s.adminHandler.DeletePost)

	// رد کردن تمام روت‌ها از داخل میدل‌ور لاگر
	return middleware.RequestLogger(mux)
}
