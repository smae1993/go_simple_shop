package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "file:app.db?_pragma=journal_mode(WAL)")
	if err != nil {
		log.Fatalf("error opening SQLite file: %v", err)
	}
	defer db.Close()

	schemaSQL, err := os.ReadFile("db/migrations/001_schema.sql")
	if err != nil {
		log.Fatalf("migration file not found: %v", err)
	}
	if _, err := db.Exec(string(schemaSQL)); err != nil {
		log.Fatalf("error creating database table: %v", err)
	}

	// ۳. تزریق وابستگی‌ها (Dependency Injection)
	// queries := database.New(db)
	// userService := service.NewUserService(queries)
	// userHandler := handler.NewUserHandler(userService)

	// ۴. تعریف روت‌ها
	mux := http.NewServeMux()
	// mux.HandleFunc("GET /", userHandler.Index)
	// mux.HandleFunc("POST /users", userHandler.Store)
	// mux.HandleFunc("GET /users/{id}/edit", userHandler.EditForm)
	// mux.HandleFunc("POST /users/{id}/update", userHandler.Update)
	// mux.HandleFunc("POST /users/{id}/delete", userHandler.Delete)

	// ۵. روشن کردن سرور
	port := ":5050"
	fmt.Printf("🚀 Server started on http://localhost%s\n", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatal(err)
	}
}
