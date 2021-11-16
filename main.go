package main

import (
	"github.com/ngocchien/presearch-tool/biz"
	"github.com/ngocchien/presearch-tool/constant"
	"github.com/ngocchien/presearch-tool/master"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
}

func main() {
	log.Info("Run App")
	m := master.NewMaster(constant.ApiUrl)
	b := biz.NewBiz(m)
	b.RunProcess()
	log.Info("Done all task! Stop!")
}
