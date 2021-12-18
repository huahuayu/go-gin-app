package session

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/huahuayu/go-gin-app/common/redis"
	"github.com/huahuayu/go-gin-app/global"
	"github.com/huahuayu/go-gin-app/model"
)

func Set(sid string, user *model.TUser) {
	key := fmt.Sprintf(redis.KEY_USER_SESSION, sid)
	data, _ := json.Marshal(user)
	redis.Cli.Set(context.Background(), key, data, global.SessionExpiredTime)
}

func Get(sid string) *model.TUser {
	key := fmt.Sprintf(redis.KEY_USER_SESSION, sid)
	res := redis.Cli.Get(context.Background(), key)
	user := &model.TUser{}
	bytes, _ := res.Bytes()
	err := json.Unmarshal(bytes, user)
	if err != nil {
		return nil
	}
	return user
}

func Del(sid string) {
	key := fmt.Sprintf(redis.KEY_USER_SESSION, sid)
	redis.Cli.Del(context.Background(), key)
}
