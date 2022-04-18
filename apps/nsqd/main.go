package main

import (
	"log"
	"nsqtoy/nsqd"
	"nsqtoy/service"
	"syscall"
)

type Program struct {
	nsqd *nsqd.NSQD
}

func (p *Program) Init() error  {
	//read option
	options := &nsqd.Options{}
	nsqd, err  := nsqd.New(options)
	if err != nil {
		return err
	}
	p.nsqd = nsqd
	return nil
}

func (p *Program) Start() error  {
	return nil
}

func (p *Program) Stop() error{
	return nil
}

func main() {
	svc := &Program{}
	if err := service.Run(svc, syscall.SIGINT, syscall.SIGTERM); err != nil {
		log.Fatal(err)
	}
}
