package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/padhitheboss/sangeet/pkg/config"
	"github.com/padhitheboss/sangeet/pkg/controller"
)

func main() {
	fmt.Println("Starting Sangeet Server")
	config.DiscordService.AddHandler(controller.MessageCreate)
	err := config.DiscordService.Open()
	if err != nil {
		log.Fatal("Error connecting to discord server ", err)
	}
	exitChan := make(chan os.Signal, 1)
	signal.Notify(exitChan, syscall.SIGINT, syscall.SIGTERM)

	// Block until a termination signal is received
	<-exitChan
	fmt.Println("Sangeet Server is shutting down...")
	// Perform any cleanup or shutdown tasks here
	// time.Sleep(1 * time.Second)
	fmt.Println("Sangeet Server has exited.")
}
