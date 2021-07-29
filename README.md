### 使用说明

一款定时企业微信定时推送工具，支持crontab语法规则。

```bash
go get github.com/pingyeaa/cowechat-notice
cowechat-notice -c [配置文件]
```

配置文件默认为`config.json`

### 配置文件示例

```json
{
  "config": [
    {
      "message": "请大家及时提交日报。@all",
      "webhook": "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=289-1a97-4994-9044-aac3ac6f5f44",
      "crontab": "*/5 * * * *"
    },
    {
      "message": "123123",
      "webhook": "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=7289-1a97-4994-9044-aac3ac6f5f44",
      "crontab": "@every 10s"
    }
  ]
}
```