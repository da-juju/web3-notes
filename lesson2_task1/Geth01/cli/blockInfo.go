package cli

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"math/big"
)

// BlockHeaderInfo 获取区块头
// 调用客户端的 HeaderByNumber 来返回有关一个区块的头信息。若传入 nil，它将返回最新的区块头。
func BlockHeaderInfo() {
	client := GetClient()
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Number: %v\n", header.Number)
}

// BlockHeaderInfoByNumber 获取区块信息
// 调用客户端的 BlockByNumber 来返回一个区块的详细信息。若传入 nil，它将返回最新的区块。
func BlockHeaderInfoByNumber() {
	// 指定区块号
	blockNumber := big.NewInt(5671744)
	client := GetClient()
	header, err := client.HeaderByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(header.Number.Uint64())     // 区块的序号，表示它在区块链中的位置：5671744
	fmt.Println(header.Time)                // 区块的时间戳，表示区块被创建的时间：1712798400
	fmt.Println(header.Difficulty.Uint64()) // 区块的难度值，用于挖矿过程中的工作量证明：0
	fmt.Println(header.Hash().Hex())        // 与Nonce一起用于工作量证明的混合哈希值：0x136d34eca31d5952021e8244b59d9c6a5deca76d1e086dbc35d024d19775cfac
}

// BlockInfo 获取区块完整信息
// 调用客户端的 BlockByNumber 方法来获得完整区块。
// 您可以读取该区块的所有内容和元数据，例如，区块号，区块时间戳，区块摘要，区块难度以及交易列表等等。
func BlockInfo() {
	blockNumber := big.NewInt(5671744)
	client := GetClient()
	// 	type Block struct {
	//	header       *Header      // 区块头
	//	uncles       []*Header	  // 区块叔块
	//	transactions Transactions // 交易列表
	//
	//	// caches
	//	hash atomic.Value
	//	size atomic.Value
	//
	//	// Td is used by package core to store the total difficulty
	//	// of the chain up to and including the block.
	//	td *big.Int               // 存储到达并包括当前区块的链的总难
	//
	//	// These fields are used by package eth to track
	//	// inter-peer block relay.
	//	ReceivedAt   time.Time    // 该区块被接收的时间
	//	ReceivedFrom interface{}  // 该区块的来源
	//}
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(block.Number().Uint64())     // 5671744
	fmt.Println(block.Time())                // 1712798400
	fmt.Println(block.Difficulty().Uint64()) // 0
	fmt.Println(block.Hash().Hex())          // 0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5
	fmt.Println(len(block.Transactions()))   // block.Transactions()返回一个包含区块中所有交易的切片，len()函数计算切片中的元素数量，即交易的数量：70
}

// TransactionCount 获取区块中交易的数量
func TransactionCount() {
	// 指定区块哈希
	blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5") // 替换为实际的区块哈希
	client := GetClient()
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count) // 交易的数量：70
}

// TX 查询区块的交易信息
func TX() {
	client := GetClient()
	block, err := client.BlockByNumber(context.Background(), big.NewInt(5671744))
	if err != nil {
		log.Fatal(err)
	}

	// 获取区块的交易信息
	//// type Transactions []*Transaction
	transactions := block.Transactions()
	fmt.Printf("transactions len：%d\n", len(transactions))
	for _, tx := range transactions {
		fmt.Println(tx.Hash().Hex())        // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
		fmt.Println(tx.Value().String())    // 交易转移的以太币数量，这里是以太币的最小单位“wei”表示的：100000000000000000
		fmt.Println(tx.Gas())               // 21000
		fmt.Println(tx.GasPrice().Uint64()) // 100000000000
		fmt.Println(tx.Nonce())             // 交易的nonce值，这是一个序列号，用于防止重放攻击，确保交易按顺序执行：245132
		fmt.Println(tx.Data())              // 交易的数据字段，对于普通的以太币转账交易，这个字段通常是空的：[]
		fmt.Println(tx.To().Hex())          // 交易的接收者地址，这是一个以太坊账户地址，以十六进制格式显示：0x8F9aFd209339088Ced7Bc0f57Fe08566ADda3587

		// ChainID()：是一个方法用于获取当前连接的以太坊网络的链ID。
		// 链ID是以太坊网络的一个标识符，不同的网络（比如主网、测试网）有不同的链ID。
		chainID, _ := client.ChainID(context.Background())
		// 使用EIP155签名者(NewEIP155Signer)，创建一个新的签名者实例，这个签名者会根据链ID来正确地签名交易
		// types.Sender函数用于从交易中恢复发送者的地址
		// 如果没有错误发生，打印出发送者的地址（以十六进制格式）
		if sender, err := types.Sender(types.NewEIP155Signer(chainID), tx); err == nil {
			fmt.Println("sender", sender.Hex()) // 0x2CdA41645F2dBffB852a605E92B185501801FC28
		}

		//if msg, err := tx.AsMessage(types.NewEIP155Signer(chainID)); err == nil {
		//	fmt.Println(msg.From().Hex()) // 0x0fD081e3Bb178dc45c0cb23202069ddA57064258
		//}

		// 通过交易的哈希值获取交易的收据信息
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		// 交易收据的状态：状态为1 ->表示交易成功被区块链网络接受
		fmt.Println(receipt.Status) // 1
		break
	}
}

// Tx2 在不获取块的情况下遍历倨交易的另一种方式是调用客户端的 TransactionInBlock 方法。
func Tx2() {
	client := GetClient()
	// 指定区块哈希
	blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5") // 替换为实际的区块哈希
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}

	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(tx.Hash().Hex()) // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
	}
}

// Tx3 还可以使用 TransactionByHash 在给定具体交易哈希值的情况下直接查询单个事务。
func Tx3() {
	// 替换为实际的交易哈希
	txHash := common.HexToHash("0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5")
	client := GetClient()
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(tx.Hash().Hex()) // 0x5d49fcaa394c97ec8a9c3e7bd9e8388d420fb050a52083ca52ff24b3b65bc9c2
	fmt.Println(isPending)       // false
}

/*
	TODO 交易和收据的理解：这两个操作都涉及到同一个区块的数据，但是它们获取的信息类型不同。
	① transactions 获取的是交易记录，
	② 而 BlockReceipts 获取的是这些交易的执行结果（接收收据）。
*/

// Receipts 查询交易收据(交易执行结果)
func Receipts() {
	blockNumber := big.NewInt(5671744)
	blockHash := common.HexToHash("0xae713dea1419ac72b928ebe6ba9915cd4fc1ef125a606f90f5e783c47cb1a4b5") // 替换为实际的区块哈希

	client := GetClient()
	receiptByHash, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithHash(blockHash, false))
	if err != nil {
		log.Fatal(err)
	}

	receiptsByNum, err := client.BlockReceipts(context.Background(), rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(blockNumber.Int64())))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(receiptByHash[0] == receiptsByNum[0]) // true
	fmt.Printf("receiptByHash len ：%d\n", len(receiptByHash))

	for _, receipt := range receiptByHash {
		fmt.Println(receipt.Status)           // 1
		fmt.Println(receipt.Logs)             // []
		fmt.Println(receipt.TxHash.Hex())     // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
		fmt.Println(receipt.TransactionIndex) // 0
		break
	}

	// 根据交易Hash获取交易收据
	txHash := common.HexToHash("0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5")
	receipt, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(receipt.Status)                // 1
	fmt.Println(receipt.Logs)                  // []
	fmt.Println(receipt.TxHash.Hex())          // 0x20294a03e8766e9aeab58327fc4112756017c6c28f6f99c7722f4a29075601c5
	fmt.Println(receipt.TransactionIndex)      // 0
	fmt.Println(receipt.ContractAddress.Hex()) // 0x0000000000000000000000000000000000000000
}

// SubBlock 订阅区块
func SubBlock() {
	// 订阅区块需要 websocket RPC URL。
	const wssUrl = "wss://ethereum-sepolia-rpc.publicnode.com"
	client, err := ethclient.Dial(wssUrl)
	if err != nil {
		log.Fatal(err)
	}

	// 创建一个新的通道，用于接收最新的区块头。
	headers := make(chan *types.Header)

	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	// 使用一个 select 语句来监听新消息。
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println(block.Hash().Hex())        // 0x99bb0f938f642b1cf7bd37964abe61f4ce15ad64421146c5edf06cee59a8cc27
			fmt.Println(block.Number().Uint64())   // 8178427
			fmt.Println(block.Time())              // 1745401836
			fmt.Println(block.Nonce())             // 0
			fmt.Println(len(block.Transactions())) // 145
		}
	}
}
