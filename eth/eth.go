package eth

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/huahuayu/go-gin-app/common/config"
	"github.com/sirupsen/logrus"
	"math/big"
)

var (
	PrivateKey     *ecdsa.PrivateKey
	PublicKeyECDSA *ecdsa.PublicKey
	EthClient      *ethclient.Client
	RpcClient      *rpc.Client
	ChainId        *big.Int
	FromAddress    common.Address
	Nonce          uint64
)

func Init() {
	var err error
	if PrivateKey, err = crypto.HexToECDSA(config.App.Eth.PrivateKey); err != nil {
		logrus.Fatal("invalid private key")
	} else {
		PublicKeyECDSA, _ = PrivateKey.Public().(*ecdsa.PublicKey)
	}
	if EthClient, err = ethclient.DialContext(context.Background(), config.App.Eth.Node); err != nil {
		logrus.Fatal("ethClient init: ", err)
	}
	if RpcClient, err = rpc.Dial(config.App.Eth.Node); err != nil {
		logrus.Fatal("rpcClient init: ", err)
	}
	if ChainId, err = EthClient.ChainID(context.Background()); err != nil {
		logrus.Fatal("get chainId error: ", err)
	}
	FromAddress = crypto.PubkeyToAddress(*PublicKeyECDSA)
	if Nonce, err = EthClient.PendingNonceAt(context.Background(), FromAddress); err != nil {
		logrus.Fatal("get nonce error: ", err)
	}
}
