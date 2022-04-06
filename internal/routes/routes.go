package routes

import (
	"context"
	"github.com/felixge/httpsnoop"
	"github.com/gorilla/mux"
	"github.com/nkralles/masters-web/internal/logger"
	"net"
	"net/http"
	"path"
	"regexp"
	"strings"
	"time"
)

func ListenAndServe(ctx context.Context, addr string) error {
	var handler = NewRouter()
	//handler = handlers.LoggingHandler(os.Stdout, handler)
	handler = LoggingHandler(handler)

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

func LoggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m := httpsnoop.CaptureMetrics(next, w, r)
		go func(r *http.Request) {
			ip, _, err := net.SplitHostPort(r.RemoteAddr)
			if err != nil {
				logger.REST.Warnf("uperip: %q is not IP:Port", r.RemoteAddr)
				return
			}
			userIP := net.ParseIP(ip)
			if userIP == nil {
				logger.REST.Warnf("uperip: %q is not IP:Port", r.RemoteAddr)
				return
			}
			logger.REST.Debug(userIP.String())
		}(r)
		logger.REST.Infof("%s %s %s %s %d %d %s", r.RemoteAddr, r.Method, r.URL.Path, r.Proto, m.Code, m.Written, m.Duration.String())
	})
}

func NewRouter() http.Handler {
	router := mux.NewRouter()
	fileServer := http.FileServer(http.Dir("build"))
	pagePathRe := regexp.MustCompile(`^/[a-zA-Z0-9/_-]+$\??.*`)

	router.Use(corsMiddleware)

	api := router.PathPrefix("/api").Subrouter()
	api.StrictSlash(true)

	api.HandleFunc("/entries", GetEntries).Methods(http.MethodGet)
	api.HandleFunc("/entries.csv", GetEntriesCSV).Methods(http.MethodGet)
	api.HandleFunc("/entries.html", GetEntriesHtml).Methods(http.MethodGet)

	api.HandleFunc("/golfers", GetGolfers).Methods(http.MethodGet)
	api.HandleFunc("/golfers/{player_id}", GetGolfer).Methods(http.MethodGet)

	api.HandleFunc("/holes", GetHoles).Methods(http.MethodGet)
	api.HandleFunc("/holes/{hole}", GetHole).Methods(http.MethodGet)

	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api/") {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}
		h := r.Header.Get("Accept")
		if strings.Contains(h, "text/html") || strings.Contains(h, "text/*") || strings.Contains(h, "*/*") {
			if r.URL.Path == "/" || pagePathRe.MatchString(r.URL.Path) {
				http.ServeFile(w, r, path.Join("build", "index.html"))
				return
			}
		}
		fileServer.ServeHTTP(w, r)
	})

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
