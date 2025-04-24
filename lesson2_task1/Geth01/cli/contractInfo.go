package cli

import (
	"Geth01/contract"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"
)

// Deploy 部署合约
func Deploy() {
	client := GetClient()

	// 创建私钥
	privateKey, err := crypto.HexToECDSA("your-private-key")
	if err != nil {
		log.Fatalf("Failed to create private key: %v", err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress) // 获取账户nonce
	if err != nil {
		log.Fatalf("Failed to retrieve account nonce: %v", err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	chainId, err := client.NetworkID(context.Background())

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address, tx, instance, err := contract.DeployStroage(auth, client) // 部署合约
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())   // 合约地址   0xcF1AFc5A88F73Da53f28dBD874D7ca14995B3CFB
	fmt.Println(tx.Hash().Hex()) // 交易哈希   0xcc3e74b9dba566f1846d5c8b9277c9a9133d9c78f7a1805e52748632b10bb556
	_ = instance                 // 合约实例
}

// CreateContractInstance 创建合约实例（加载合约）
func CreateContractInstance() *contract.Stroage {
	client := GetClient()
	const contractAddr = "0xcF1AFc5A88F73Da53f28dBD874D7ca14995B3CFB"
	contractInstance, err := contract.NewStroage(common.HexToAddress(contractAddr), client)
	if err != nil {
		log.Fatal(err)
	}
	return contractInstance
}

// CallContract 执行合约
func CallContract() {
	instance := CreateContractInstance()
	// 初始化交易opt实例
	privateKey, err := crypto.HexToECDSA("your-private-key")
	if err != nil {
		log.Fatalf("Failed to create private key: %v", err)
	}
	chainID, _ := GetClient().ChainID(context.Background())
	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	tx, err := instance.Store(opt, big.NewInt(100))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx.Hash().Hex()) // 交易哈希

	callOpt := &bind.CallOpts{Context: context.Background()}
	val, err := instance.Retrieve(callOpt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(val) // 100
}

// QueryContractEvent 查询合约事件
func QueryContractEvent() {
	client := GetClient()

	const contractAddr = "0xcF1AFc5A88F73Da53f28dBD874D7ca14995B3CFB"
	const StorageABI = `[{"anonymous":false,"inputs":[{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"}],"name":"MyEvent","type":"event"},{"inputs":[],"name":"number","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"retrieve","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"num","type":"uint256"}],"name":"store","outputs":[],"stateMutability":"nonpayable","type":"function"}]`

	// 构造查询过滤器，用于筛选特定合约地址的事件日志
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(8183426),
		// ToBlock:   big.NewInt(2394201),
		Addresses: []common.Address{
			common.HexToAddress(contractAddr),
		},
		// Topics: [][]common.Hash{
		//  {},
		//  {},
		// },
	}

	// 使用过滤器查询事件日志
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	// 解析ABI编码信息
	contractAbi, err := abi.JSON(strings.NewReader(StorageABI))
	if err != nil {
		log.Fatal(err)
	}

	// 遍历查询到的事件日志
	for _, vLog := range logs {
		fmt.Println(vLog.BlockHash.Hex())
		fmt.Println(vLog.BlockNumber)
		fmt.Println(vLog.TxHash.Hex())

		// 定义一个匿名结构体用于接收解包后的事件数据
		// 定义了一个匿名结构体，其中包含一个名为Value的字段，该字段的类型为*big.Int（指向big.Int类型的指针）
		event := struct {
			Value *big.Int `json:"value"`
		}{} // 这里的{}表示初始化这个匿名结构体的实例，由于没有提供Value字段的值，所以它的零值将是nil

		// 将事件日志中的数据解包到event结构体中
		err := contractAbi.UnpackIntoInterface(&event, "MyEvent", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(event.Value) // 打印事件中的Value值

		// indexed topics，被索引的事件信息，存放在topics的区域
		//// 处理被索引的事件主题（topics）
		var topics []string
		for i := range vLog.Topics {
			topics = append(topics, vLog.Topics[i].Hex()) // 将每个主题转换为十六进制字符串，追加到topics切片中
		}
		// 打印第一个主题（通常是事件签名）
		fmt.Println("topics[0]=", topics[0]) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
		// 如果有额外的索引主题，则打印它们
		if len(topics) > 1 {
			fmt.Println("indexed topics:", topics[1:])
		}
	}
}

// SubContractEvent 订阅合约事件
func SubContractEvent() {
	// 订阅合约事件需要 websocket RPC URL。
	const wssUrl = "wss://ethereum-sepolia-rpc.publicnode.com"
	const StorageABI = `[{"anonymous":false,"inputs":[{"indexed":false,"internalType":"uint256","name":"value","type":"uint256"}],"name":"MyEvent","type":"event"},{"inputs":[],"name":"number","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"retrieve","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"num","type":"uint256"}],"name":"store","outputs":[],"stateMutability":"nonpayable","type":"function"}]`

	client, err := ethclient.Dial(wssUrl)
	if err != nil {
		log.Fatal(err)
	}

	// 构造查询过滤器，用于筛选特定合约地址的事件日志
	contractAddress := common.HexToAddress("0xcF1AFc5A88F73Da53f28dBD874D7ca14995B3CFB")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	// 创建一个新的通道，用于接收最新的事件日志。
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	// 解析ABI编码信息
	contractAbi, err := abi.JSON(strings.NewReader(StorageABI))
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println("sub:" + vLog.BlockHash.Hex())
			fmt.Println("sub:", vLog.BlockNumber)
			fmt.Println("sub:" + vLog.TxHash.Hex())

			// 定义一个匿名结构体用于接收解包后的事件数据
			// 定义了一个匿名结构体，其中包含一个名为Value的字段，该字段的类型为*big.Int（指向big.Int类型的指针）
			event := struct {
				Value *big.Int `json:"value"`
			}{} // 这里的{}表示初始化这个匿名结构体的实例，由于没有提供Value字段的值，所以它的零值将是nil
			// 将事件日志中的数据解包到event结构体中
			err := contractAbi.UnpackIntoInterface(&event, "MyEvent", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(event.Value) // 打印事件中的Value值

			var topics []string
			for i := range vLog.Topics {
				topics = append(topics, vLog.Topics[i].Hex())
			}
			fmt.Println("topics[0]=", topics[0])
			if len(topics) > 1 {
				fmt.Println("index topic:", topics[1:])
			}
		}
	}
}
