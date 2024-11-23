package main

import (
	"github.com/maksimmironenko/cloud/internal/worker"
)

func main() {
	worker.RunWorkers(3)
}
