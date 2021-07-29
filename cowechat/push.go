package cowechat

import (
	"fmt"
	"io/ioutil"

	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// Push 推送消息
func Push(webHook string, content string) (bool, error) {
	client := ghttp.NewClient()
	client.SetHeader("Content-Type", "application/json")
	data := g.Map{
		"msgtype": "markdown",
		"markdown": g.Map{
			"content": content,
		},
	}
	res, err := client.Post(webHook, data)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}
