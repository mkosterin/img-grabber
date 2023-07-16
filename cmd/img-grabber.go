package main

import (
	"img-grabber/internal/pkg/app"
	"log"
	"os"
	"time"
)

func main() {

	result := make(chan bool, 1)

	go job(result)

	select {
	case <-result:
		os.Exit(0)
	case <-time.After(10 * time.Second):
		log.Println("Timeout has been reached")
	}

}

func job(result chan<- bool) {
	a := app.New()
	err := a.Run()
	if err != nil {
		log.Fatal(err)
	}
	result <- true
}
