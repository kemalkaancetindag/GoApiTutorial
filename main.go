package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"mongoapp.com/config"
	"mongoapp.com/handlers"
	"mongoapp.com/repository"
	"mongoapp.com/services"
)

func main() {
	l := log.New(os.Stdout, "todo-api", log.LstdFlags)
	config.ConnectDB()
	dbClient := config.GetCollection(config.DB, "Todos")
	todoRepo := repository.NewTodoRepositoryDB(dbClient)
	todoService := services.NewTodoService(&todoRepo)

	th := handlers.NewTodoHandler(todoService)
	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/todo", th.GetAll)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/todo", th.Insert)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	s.Shutdown(tc)
}
