package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"fmt"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/santoshdeshpande/realworld/pkg/api/users"
	"github.com/santoshdeshpande/realworld/pkg/mw"
)

//Server is a http server that serves the realworld endpoints
type Server struct {
	srv    *http.Server
	logger *zap.Logger
}

//NewServer returns a new server
func NewServer(port int) (*Server, error) {
	logger, err := zap.NewDevelopment(zap.AddStacktrace(zapcore.FatalLevel))
	if err != nil {
		log.Fatalf("Can't initialize zap logger: %v", err)
	}
	defer logger.Sync()

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(mw.Logger(logger))
	r.Get("/", indexHandler)
	// us := &users.UserService{}
	ur := users.NewUserResource(logger)
	r.Mount("/users", ur.Routes())
	address := fmt.Sprintf(":%d", port)

	errorLog, _ := zap.NewStdLogAt(logger, zap.ErrorLevel)
	server := &http.Server{
		Addr:         address,
		Handler:      r,
		ErrorLog:     errorLog,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	httpServer := &Server{logger: logger, srv: server}

	return httpServer, nil
}

//Start the server
func (s *Server) Start() {
	sugar := s.logger.Sugar()
	done := make(chan bool, 1)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		sugar.Info("Server is shutting down...")
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		s.srv.SetKeepAlivesEnabled(false)
		if err := s.srv.Shutdown(ctx); err != nil {
			sugar.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		}
		close(done)
	}()

	url := fmt.Sprintf("http://localhost%s", s.srv.Addr)
	sugar.Infow("Server is ready to handle requests at", "port", s.srv.Addr, "url", url)

	if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		sugar.Fatalf("Could not listen on %s: %v\n", s.srv.Addr, err)
	}

	<-done
	sugar.Info("Server stopped")

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
