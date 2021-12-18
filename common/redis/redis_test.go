package redis

import (
	"context"
	"github.com/huahuayu/go-gin-app/common/config"
	"testing"
)

func TestClient_SubscribeA(t *testing.T) {
	config.Init("", "dev")
	Init()
	pubSub := Cli.Subscribe(context.Background(), "ETH:NewTxns|All")
	_, err := pubSub.Receive(context.Background())
	if err != nil {
		panic(err)
	}
	for msg := range pubSub.Channel() {
		t.Log("test sub: ", msg.Channel, msg.Payload)
	}
}

func TestClient_SubscribeB(t *testing.T) {
	config.Init("", "dev")
	Init()
	pubSub := Cli.Subscribe(context.Background(), "BSC:NewTxns|All")
	_, err := pubSub.Receive(context.Background())
	if err != nil {
		panic(err)
	}
	for msg := range pubSub.Channel() {
		t.Log("test sub: ", msg.Channel, msg.Payload)
	}
}
