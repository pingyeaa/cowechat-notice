package main

import (
	"cowechat-notice/config"
	"cowechat-notice/cowechat"
	"cowechat-notice/cron"
)

func main() {
	confs, err := config.GetConfig()
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
