package redis

import "github.com/huahuayu/go-gin-app/common/config"

var Cli Client

func Init() {
	var err error
	Cli, err = NewClient(config.App.Redis.Host, config.App.Redis.Pass, config.App.Redis.Db)
	if err != nil {
		panic(err)
	}
}
