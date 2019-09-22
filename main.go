package main

import (
	"fmt"
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
	fmt.Println(Storage)
}
