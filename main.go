package main

import (
	"github.com/evsio0n/log"
	"netflix_agent/utils"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var err error

func init() {
	log.SetDebug(utils.IsDebug)
	log.IsShowLogCatagory(false)
	log.SetSyslog(utils.IsSyslog, "netflix-agent")
	log.IsShowDate(utils.LogHaseDate)
	err = utils.ChangeIPv6()
	if err != nil {
		log.Fatal(err.Error())
		os.Exit(1)
	}
}

func main() {
	signalChannel := make(chan os.Signal, 2)
	signal.Notify(signalChannel, os.Interrupt, syscall.SIGTERM)
	go func() {
		sig := <-signalChannel
		switch sig {
		case os.Interrupt:
			//handle SIGINT
			log.Info("SIGINT received. Shutting down...")
			utils.FlushNAT()
		case syscall.SIGTERM:
			//handle SIGTERM
			log.Info("SIGINT received. Shutting down...")
			utils.FlushNAT()
		}
	}()
	detectBlock()
}

func detectBlock() {
	for true {
		if !utils.TestUnblock() {
			log.Warn("Unblock failing, changing IP...")
			err := utils.ChangeIPv6()
			if err != nil {
				log.Error(err.Error())
				os.Exit(1)
			}
		} else {
			log.Info("Unblock OK")
		}
		time.Sleep(time.Second * 30)
	}
}
