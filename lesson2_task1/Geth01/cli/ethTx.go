package cli

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
)

// EthTx ETH转账
func EthTx() {
	client := GetClient()

	// 将十六进制的私钥字符串转换为ECDSA私钥对象。如果转换失败，记录错误并退出程序
	//privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
	privateKey, err := crypto.HexToECDSA("1cfbd8cb8165a85eae14834992524bdf8cb058fb4d32f814ee3f0fa565e1e775")
	if err != nil {
		log.Fatal(err)
	}

	// 从私钥生成公钥，并从公钥生成以太坊地址。
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA) // 转账发起人地址
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(50000000000000000) // in wei (0.05 eth)
	gasLimit := uint64(21000)              // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d") // 转账接受人地址
	toAddress := common.HexToAddress("0xe2BB7Af706d51fABa2d177568fCc51C6f8Fe692B") // 转账接受人地址
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex()) // 0xd5a0c5d864c761a56924d0d26a42531a8006e4b10805517367aaa63e7624d705
}

// BalanceAt 获取账户余额
// 调用 ethclient 的 BalanceAt 方法，给它传递账户地址和可选的区块号。将区块号设置为 nil 将返回最新的余额
func BalanceAt() {
	client := GetClient()
	accountAddress := common.HexToAddress("0xA9F76c1f99FF3F080d303775AB9615CE43271A9C")
	balanceAt, err := client.BalanceAt(context.Background(), accountAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("balance: %v\n", balanceAt) // 单位wei
	// 创建一个big.Float类型的变量
	floatNum := new(big.Float)
	// 将big.Int转换为big.Float
	floatNum.SetInt(balanceAt)
	// 使用Quo方法来进行除法运算，将floatNum除以1e18（即10的18次方）
	etherValue := new(big.Float).Quo(floatNum, big.NewFloat(1e18))
	fmt.Printf("etherValue: %v", etherValue) // 单位ether
}
