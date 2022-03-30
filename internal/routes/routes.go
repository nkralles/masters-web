package routes

import (
	"context"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"net"
	"net/http"
	"os"
	"time"
)

func ListenAndServe(ctx context.Context, addr string) error {
	var handler = NewRouter()
	handler = handlers.LoggingHandler(os.Stdout, handler)

	srv := &http.Server{
		Handler:      handler,
		Addr:         addr,
		WriteTimeout: 10 * time.Minute,
		ReadTimeout:  60 * time.Second,
		BaseContext: func(l net.Listener) context.Context {
			return ctx
		},
	}
	go func() {
		<-ctx.Done()
		sctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		srv.Shutdown(sctx)
	}()

	return srv.ListenAndServe()
}

func NewRouter() http.Handler {
	router := mux.NewRouter()
	router.Use(corsMiddleware)

	api := router.PathPrefix("/api").Subrouter()
	api.StrictSlash(true)

	api.HandleFunc("/holes", GetHoles).Methods(http.MethodGet)

	return router
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
		w.Header().Set("Access-Control-Allow-Method", "POST, GET, DELETE, PUT, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Max-Age", "86400")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
