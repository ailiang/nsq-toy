package service

import (
	"os"
	"os/signal"
	"syscall"
)

type Service interface {
	Init() error
	Start() error
	Stop() error
}

func Run(s Service, sig ...os.Signal) error {
	if len(sig) == 0 {
		sig = []os.Signal{syscall.SIGINT}
	}

	if err := s.Init() ; err != nil {
		return  err
	}
	if err := s.Start(); err != nil {
		return err
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, sig...)
	select {
	case <-sigChan:
	}
	return  s.Stop()
}

