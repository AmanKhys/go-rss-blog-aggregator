package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	//	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
  fmt.Println("hello world")

  godotenv.Load(".env")

  portString := os.Getenv("PORT")
  if portString == "" {
    log.Fatal("PORT is not bound in the enviorument")
  }
  fmt.Println("PORT:", portString)
  
  router := chi.NewRouter()

  router.Use(cors.Handler(cors.Options{
    AllowedOrigins: []string{"https://*"},
    AllowedMethods: []string{"GET", "POST", "DELETE", "OPTIONS"},
    AllowedHeaders: []string{"*"},
    ExposedHeaders: []string{"Link"},
    AllowCredentials: false,
    MaxAge: 300,
  }))

  srv := &http.Server{
    Handler: router,
    Addr: ":" + portString,
  }

  log.Printf("Server starting on port %v", portString)
  err := srv.ListenAndServe()
  if err != nil {
    log.Fatal(err)
  }

  

}
