package home

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

const message = "Hello Servian"

type HomeModule struct {
	logger *log.Logger
}

func (h *HomeModule) HomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain:charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

func (h *HomeModule) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		starttime := time.Now()
		defer h.logger.Printf("Request processed in %s", time.Now().Sub(starttime))
		next(w, r)

	}
}

func (h *HomeModule) SetupRoutes(mux *mux.Router) {
	mux.HandleFunc("/", h.HomePage)
	mux.HandleFunc("/home", h.HomePage)
}

func NewHomeModule(log *log.Logger) *HomeModule {
	return &HomeModule{
		logger: log,
	}
}
