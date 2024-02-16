package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	r := chi.NewRouter()

	godotenv.Load()
	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("port not found in env")
	}

	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

  v1Router := chi.NewRouter()

  v1Router.Get("/healthz", handlerReadiness)
  v1Router.Get("/err",handlerErr)

  r.Mount("/v1", v1Router)

	fmt.Println("listening on port:" + portString)

	server := &http.Server{
		Handler: r,
		Addr:    ":" + portString,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
