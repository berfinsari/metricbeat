package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/berfinsari/metricbeat/membeat/beater"
)

func main() {
	err := beat.Run("membeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
