package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/techievee/godemo/home"
	"github.com/techievee/godemo/server"
	"github.com/techievee/godemo/storage"
	"github.com/techievee/godemo/user"
	"log"
	"os"
)

const message = "Servian"

//Set of variables that need to be fetched from the OS Env or Kubernetes config map/Secrets
var (
	Storage        = os.Getenv("STORAGE")
	StorageDB      = os.Getenv("STORAGE_DB")
	ServiceAddress = os.Getenv("SERVICE_ADDRESS")
	RedisPassword  = os.Getenv("REDIS_PASS")
)

func main() {
	fmt.Println("Program Started")
	//Create a custom logger with standard specification
	logger := log.New(os.Stdout, "GO_LOG:", log.LstdFlags|log.Lshortfile)

	//------------Backend Init and check--------------------
	var backend *storage.RedisClient
	backend = storage.NewRedisClient(Storage, RedisPassword, StorageDB)
	pong, err := backend.Check()
	if err != nil {
		log.Printf("Can't establish connection to backend: %v", err)
		os.Exit(0)
	} else {
		log.Printf("Backend check reply: %v", pong)
	}
	//----------------------------------------------------

	//Initizlize the modules
	h := home.NewHomeModule(logger)
	u := user.NewUserModule(logger, backend)

	//Create a mux, to serve the Endpoints
	router := mux.NewRouter()
	h.SetupRoutes(router)
	u.SetupRoutes(router)

	//Create server and serve
	srv := server.NewServer(router, ServiceAddress)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Server not started : %v", err)
	}

}
