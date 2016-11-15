package main

import (
	"errors"
	"github.com/gwhn/log-factory/logger"
)

func main() {
	log := logger.New()
	log.Debug("%d + %d = %d", 2, 3, 5)
	log.Warning("watch out!")
	log.Info("i'm here")
	log.Error("something's gone wrong with: %v", errors.New("ME"))
	log.Critical("flat line, no response with: %v", errors.New("YOU"))
}