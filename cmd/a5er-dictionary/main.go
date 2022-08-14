package main

import (
	"log"

	"github.com/future-architect/a5er-dictionary/internal/cmd"
)

func main() {
	if err := cmd.Do(); err != nil {
		log.Fatalf("fail to execute: %v\n", err)
	}
}
