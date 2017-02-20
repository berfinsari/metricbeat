package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/berfinsari/lsbeat/beater"
)

func main() {
	err := beat.Run("lsbeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
