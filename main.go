package main

import (
	"fmt"
	"os"
)

type Config struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Domain   string `yaml:"domain"`
}

var AppConfig Config

func main() {
	if err := initConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "Error initializing config: %v\n", err)
		os.Exit(1)
	}

	if err := subsonicPing(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		fmt.Println("Please check your config.yaml and try again.")
		os.Exit(1)
	}

	if err := initPlayer(); err != nil {
		panic(err)
	}
	defer shutdownPlayer()
}
