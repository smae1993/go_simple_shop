package middleware

import (
	"log"
	"net/http"
	"time"
)

// میدل‌ور ثبت لاگ درخواست‌ها و زمان اجرا (معادل HTTP Kernel Middleware در لاراول)
func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("[%s] %s - %v", r.Method, r.URL.Path, time.Since(start))
	})
}
