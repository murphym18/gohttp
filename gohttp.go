package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Settings struct {
	Port    string
	Address string
	Dir     string
}

func (config *Settings) printSettings() {
	fmt.Printf("Serving '%s'\n", config.Dir)
	fmt.Printf("Listening at %s:%s\n", config.Address, config.Port)
	fmt.Printf("Press Control+C to exit\n")
}

func (config *Settings) listenAddress() string {
	return fmt.Sprintf("%s:%s", config.Address, config.Port)
}

func getWorkingDir() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func main() {
	config := Settings{Port: "8080", Address: "0.0.0.0", Dir: getWorkingDir()}
	config.printSettings()
	fs := http.FileServer(http.Dir(config.Dir))
	http.Handle("/", http.StripPrefix("/", fs))
	http.ListenAndServe(config.listenAddress(), nil)
}
