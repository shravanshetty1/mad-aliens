package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/shravanshetty1/mad-aliens/pkg/maliens"
)

func main() {
	cmd := maliens.Command()
	rand.Seed(time.Now().UnixNano())

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
