package main

import (
	"cowechat-notice/config"
	"cowechat-notice/cowechat"
	"cowechat-notice/cron"

	//_ "cowechat-notice/daemon"
	"flag"
)

func main() {

	confPath := flag.String("c", "config.json", "通过`-c`添加配置文件")
	flag.Parse()

	confs, err := config.GetConfig(*confPath)
	if err != nil {
		panic(err)
	}
	if confs == nil {
		panic(err)
	}

	c := cron.New()
	for _, conf := range confs.Config {
		_, err = c.AddFunc(conf.Crontab, func(p cron.Param) {
			_, err := cowechat.Push(p.Webhook, p.Content)
			if err != nil {
				panic(err)
			}
		}, cron.Param{
			Webhook: conf.Webhook,
			Content: conf.Message,
		})
		if err != nil {
			panic(err)
		}
	}
	c.Start()
	select {}
}
