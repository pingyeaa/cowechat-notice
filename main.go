package main

import (
	"cowechat-notice/config"
	"cowechat-notice/cowechat"
	"cowechat-notice/cron"
	"os"
)

func main() {

	confPath := "config.json"
	if len(os.Args) == 3 {
		if os.Args[1] != "-c" {
			panic("命令参考格式：cowechat-notice -c [配置文件路径]")
		}
		if os.Args[2] == "" {
			panic("命令参考格式：cowechat-notice -c [配置文件路径]")
		}
		confPath = os.Args[2]
	}

	confs, err := config.GetConfig(confPath)
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
