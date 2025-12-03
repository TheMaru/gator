package main

import (
	"fmt"
	"log"

	"github.com/TheMaru/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("reading of config not possible: %v", err)
	}

	cfg.SetUser("themaru")

	finalCfg, err := config.Read()
	if err != nil {
		log.Fatalf("second reading of config not possible: %v", err)
	}

	fmt.Println(finalCfg)
}
