package cli

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
	"log"
)

// CreateWallet 创建钱包
// 生成随机私钥 ---> 通过椭圆曲线生成公钥 --- > Keccak256压缩公钥压缩 --- > 截取后20byte生成地址
func CreateWallet() {
	// 1、需要导入 go-ethereum crypto 包，该包提供用于生成随机私钥的 GenerateKey 方法
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:]) // 去掉'0x'

	// 2、通过椭圆曲线生成公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Println("from pubKey:", hexutil.Encode(publicKeyBytes)[4:]) // 去掉'0x04'

	// 3、Keccak256压缩公钥压缩
	// 4、截取后20byte生成地址
	// 使用公钥生成以太坊地址，并将其转换为十六进制字符串打印出来
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address)

	// 创建一个新的Keccak256哈希对象，并写入公钥字节切片（去掉第一个字节，因为以太坊使用压缩公钥）
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	fmt.Println("full:", hexutil.Encode(hash.Sum(nil)[:]))
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) // 原长32位，截去12位，保留后20位
}
