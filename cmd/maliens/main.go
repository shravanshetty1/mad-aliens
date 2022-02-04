package main

import (
	"log"

	"github.com/shravanshetty1/mad-aliens/pkg/maliens"
)

func main() {
	cmd := maliens.Command()

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
