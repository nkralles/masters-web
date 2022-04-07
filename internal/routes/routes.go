package routes

import (
	"context"
	"github.com/felixge/httpsnoop"
	"github.com/gorilla/mux"
	"github.com/nkralles/masters-web/internal/logger"
	"github.com/nkralles/masters-web/internal/persistence"
	"github.com/ua-parser/uap-go/uaparser"
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
		var ip = r.Header.Get("X-Forwarded-For")
		ua := r.UserAgent()
		parser := uaparser.NewFromSaved()
		client := parser.Parse(ua)
		logger.REST.Infof("%s %s %s %s %d %d %s", ip, r.Method, r.URL.Path, r.Proto, m.Code, m.Written, m.Duration.String())
		persistence.DefaultDriver().HttpTelemetry(context.Background(), persistence.Telemetry{
			IP:           ip,
			HttpMethod:   r.Method,
			UrlPath:      r.URL.Path,
			HttpCode:     m.Code,
			HttpWritten:  m.Written,
			HttpDuration: m.Duration,
			Ua:           client,
		})
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

	api.HandleFunc("/scores", GetScores).Methods(http.MethodGet)

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
