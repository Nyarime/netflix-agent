package main

import (
	"github.com/evsio0n/log"
	"netflix_agent/utils"
	"os"
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
