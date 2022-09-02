package main

import (
	"flag"
	"log"

	"github.com/andytyc/igocase/conf"
	"github.com/andytyc/igocase/gcase"
	"github.com/andytyc/igocase/utils"
)

func init() {
	flag.IntVar(&conf.FlagCaseNum.Val, conf.FlagCaseNum.Key, conf.FlagCaseNum.ValDefault, conf.FlagCaseNum.Usage())
	flag.Parse()
	conf.FlagCheck()
}

func main() {
	gcase.Do()
	sig := <-utils.NotifySignal()
	log.Println("退出 :", sig)
}
