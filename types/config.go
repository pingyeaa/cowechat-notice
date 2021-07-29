package types

type Config struct {
	Config []struct {
		Message string `json:"message"`
		Webhook string `json:"webhook"`
		Crontab string `json:"crontab"`
	} `json:"config"`
}
