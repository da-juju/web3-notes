package cli

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

// 测试网络环境访问地址
// // Too Many Requests: {"code":-32005,"message":"Too Many Requests","data":{"see":"https://infura.io/dashboard"}}
// const clientUrl = "https://sepolia.infura.io/v3/8e1950c1a4f74c34b5741f1e9d518947"
const clientUrl = "https://ethereum-sepolia-rpc.publicnode.com"

// GetClient 获取客户端
func GetClient() *ethclient.Client {
	client, err := ethclient.Dial(clientUrl)
	defer client.Close()
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	return client
}
