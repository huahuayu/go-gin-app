package redis

import (
	"context"
	"github.com/huahuayu/go-gin-app/common/config"
	"testing"
)

func TestSubscribeA(t *testing.T) {
	config.Init("", "dev")
	Init()
	pubSub := Subscribe(context.Background(), "Ethereum:NewTxns|Uniswap|swapExactETHForTokens")
	_, err := pubSub.Receive(context.Background())
	if err != nil {
		panic(err)
	}
	for msg := range pubSub.Channel() {
		t.Log("test sub: ", msg.Channel, msg.Payload)
	}
}

func TestSubscribeB(t *testing.T) {
	config.Init("", "dev")
	Init()
	pubSub := Subscribe(context.Background(), "Ethereum:NewTxns|Uniswap|swapExactTokensForETH")
	_, err := pubSub.Receive(context.Background())
	if err != nil {
		panic(err)
	}
	for msg := range pubSub.Channel() {
		t.Log("test sub: ", msg.Channel, msg.Payload)
	}
}
