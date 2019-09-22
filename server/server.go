package server

import (
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
	"time"
)

func NewServer(router *mux.Router, serviceAddr string) *http.Server {

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, //you service is available and allowed for this base url
		AllowedMethods: []string{
			http.MethodGet, //http methods for your app
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},

		AllowedHeaders: []string{
			"*", //or you can your header key values which you are using in your application

		},
	})

	handler := corsOpts.Handler(router)

	srv := &http.Server{
		Addr:         serviceAddr,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      handler,
	}

	return srv

}
