package main

import (
	"fmt"
	"log"
	"sync"

	"github.com/intetunder-temp/learning-go/pluralsight/deep-dive-into-go-packages/working-within-a-package/services/events"
)

func init() {
	fmt.Println("init in main.go")
}

func main() {
	log.Println("Starting all services")

	// start the services
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go startServer(wg, events.StartServer)
	wg.Wait()
	log.Println("All services stopped")
}

func startServer(wg *sync.WaitGroup, startFunc func() error) {
	err := startFunc()
	wg.Done()
	if err != nil {
		log.Fatal(err)
	}
}
