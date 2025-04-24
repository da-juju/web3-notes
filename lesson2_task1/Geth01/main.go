package main

import "Geth01/cli"

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	//// 示例：创建一个新的RPC客户端
	//client, err := rpc.Dial("http://localhost:8545")
	//if err != nil {
	//	fmt.Printf("Failed to connect to the Ethereum client: %v\n", err)
	//}
	//// 关闭客户端
	//defer client.Close()
	////新建账户
	//account, err := cli.NewAccount(client, "123456")
	//if err != nil {
	//	fmt.Println("[main err]", err)
	//}
	//fmt.Println("新建账号：", account)
	//
	//// 查询所有账户
	//listAccount, err := cli.ListAccount(client)
	//if err != nil {
	//	fmt.Println("[main err]", err)
	//}
	//for inx, acc := range listAccount {
	//	fmt.Printf("账户[%d]：%s\n", inx, acc)
	//}
	//number, _ := cli.BlockNumber(client)
	//fmt.Printf("区块高度：%d\n", number)
	//
	//// -----------------------------------------------------
	//
	//// 使用Geth客户端，直接调用API(不适用cli.Call(...)方法 )
	//ethCli, _ := ethclient.Dial("http://localhost:8545")
	//defer ethCli.Close()
	//
	//// 获取区块高度
	////blockNumber, _ := ethCli.BlockByNumber(context.Background(), nil)
	//
	//// 查询以太坊账户余额
	//// common.HexToAddress：接受一个十六进制（Hex）字符串作为参数，并将其转换为一个地址类型
	//address := common.HexToAddress(listAccount[0])
	//balanceAt, _ := ethCli.BalanceAt(context.Background(), address, nil)
	//fmt.Printf("账户余额：%d\n", balanceAt)

	// -------------------ethclient使用------------------------
	//NewAccount()

	//DeployContract()

	//CallContractUpdateState()
	//// goroutine
	//wg := sync.WaitGroup{}
	//wg.Add(1)
	//go func() {
	//	EventSub()
	//	defer wg.Done()
	//}()
	//wg.Wait()

	//CallContractReadState()

	// -----------------查询区块------------------------
	// 查询最新的区块头信息
	// cli.BlockHeaderInfo()

	// 查询指定区块号的区块头信息
	//cli.BlockHeaderInfoByNumber()

	// 查询区块的完整信息
	//// 报错：transaction type not supported
	//// 处理方案：go get -u github.com/ethereum/go-ethereum
	// cli.BlockInfo()

	// 获取区块的交易数量
	//cli.TransactionCount()

	// 查询区块的交易信息
	//cli.TX()

	// 查询交易收据
	//cli.Receipts()

	// 订阅区块
	//cli.SubBlock()

	// ---------------------创建钱包-----------------------
	//cli.CreateWallet()

	// ---------------------转账-----------------------
	//cli.EthTx()

	// ---------------------查询账户余额-----------------------
	//cli.BalanceAt()

	// ----------------------合约相关-----------------------
	// 部署合约
	//cli.Deploy()

	// 调用合约
	//cli.CallContract()

	// 查询合约事件
	//cli.QueryContractEvent()

	// 订阅合约事件
	cli.SubContractEvent()
}

/*
// NewAccount 创建账户
func NewAccount() {
	//go创建以太坊账户
	//创建一个新的钱包（keystore）
	keyDir := "D:\\geth\\geth-alltools-1.9.11\\geth-alltools-windows-amd64-1.9.22-c71a7e26\\mychain\\data1\\keystore" // 指定geth私链钱包目录
	ks := keystore.NewKeyStore(keyDir, keystore.StandardScryptN, keystore.StandardScryptP)
	// 创建一个新的账户并保存到钱包
	account, err := ks.NewAccount("123") // 密码为空字符串，实际使用时应设置密码
	if err != nil {
		log.Fatalf("Failed to create account: %v", err)
	}
	fmt.Printf("Address: %s", account.Address.Hex())
}

const keyInfo = `{"address":"16a6163928b7eedf0e2cf1c9aeb7a9a1220af998","crypto":{"cipher":"aes-128-ctr","ciphertext":"10d6f5844c1aa192655618e0bf47c6a81a88f900e5f2e000be66d443e06ecfc6","cipherparams":{"iv":"f422c366bb71a8b2871a1613f562203e"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"65f06bf6135ec2505af3141656a160571a0e55fc6cfd516f7f6fdec527291a88"},"mac":"9219030cc940c45b6d6425a6f9f1a5958fe063b385c9fc7f4170ab93d636c513"},"id":"0d5f54a6-6e4c-4b2b-ace8-4cf60cf06117","version":3}`

// const clientUrl = "http://127.0.0.1:7545"   // ganache
const clientUrl = "http://127.0.0.1:8545" // geth搭建的私有链

// DeployContract 部署合约
func DeployContract() {
	// 创建客户端
	ethCli, err := ethclient.Dial(clientUrl)
	defer ethCli.Close()
	if err != nil {
		fmt.Printf("Failed to connect to the Ethereum client: %v\n", err)
	}
	// 身份认证：它将用于签名和发送交易
	// 创建了一个用于身份认证的事务执行器（transactor）。
	// 它使用之前创建的钱包，并设置了一个密码（"123"）和链ID。链ID是用于区分不同以太坊网络的一个标识符。
	chainID, _ := ethCli.ChainID(context.Background())
	fmt.Println("chainID:", chainID)
	// 创建一个用于身份认证的事务执行器（transactor）。
	// 它使用之前创建的钱包，并设置了一个密码（"123"）和链ID。链ID是用于区分不同以太坊网络的一个标识符。
	auth, err := bind.NewTransactorWithChainID(strings.NewReader(keyInfo), "123", chainID)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	// 部署合约
	//// Failed to deploy contract: insufficient funds for transfer：没有足够的资金用于部署合约
	//// 去geth客户端进行转账：eth.sendTransaction({from: address, to: address, value: web3.toWei(amount, "ether")})
	//// Failed to deploy contract: intrinsic gas too low
	//// 设置部署需要的gas：auth.GasLimit = uint64(3000000)
	auth.GasLimit = uint64(3000000)
	address, _, _, err := contract.DeployStorage(auth, ethCli)
	if err != nil {
		log.Fatalf("Failed to deploy contract: %v", err)
	}
	fmt.Println("合约地址：", address.Hex())
}

// 部署的合约地址
const conAddr = "0x35cD5cFf9415CDf080428A150d0Bc71fd6d88715"

// CallContractUpdateState 调用合约更改合约的状态
func CallContractUpdateState() {
	// 创建客户端
	ethCli, err := ethclient.Dial(clientUrl)
	defer ethCli.Close()
	if err != nil {
		fmt.Printf("Failed to connect to the Ethereum client: %v\n", err)
	}
	// 创建合约对象
	contractAddr := common.HexToAddress(conAddr) // 使用合约地址创建合约对象
	contractInstance, err := contract.NewStorage(contractAddr, ethCli)
	if err != nil {
		log.Fatalf("Failed to instantiate a Contract: %v", err)
	}
	// Create an authorized transactor and call the store function
	chainID, _ := ethCli.ChainID(context.Background())
	auth, err := bind.NewTransactorWithChainID(strings.NewReader(keyInfo), "123", chainID)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	// 调用合约方法
	auth.GasLimit = uint64(3000000) // Failed to call store: execution reverted
	tx, err := contractInstance.Store(auth, big.NewInt(100))
	if err != nil {
		log.Fatalf("Failed to call store: %v", err)
	}
	fmt.Printf("Update pending: 0x%x\n", tx.Hash())
}

// CallContractReadState 调用合约读取合约的状态
func CallContractReadState() {
	// 创建客户端
	ethCli, err := ethclient.Dial(clientUrl)
	defer ethCli.Close()
	if err != nil {
		fmt.Printf("Failed to connect to the Ethereum client: %v\n", err)
	}
	// 创建合约对象
	contractAddr := common.HexToAddress(conAddr) // 使用合约地址创建合约对象
	contractInstance, err := contract.NewStorage(contractAddr, ethCli)
	if err != nil {
		log.Fatalf("Failed to instantiate a Contract: %v", err)
	}
	// 调用合约方法
	num, err := contractInstance.Retrieve(nil)
	if err != nil {
		log.Fatalf("Failed to call retrieve: %v", err.Error())
	}
	fmt.Println("num:", num)
}

func EventSub() {
	// 获取事件订阅客户端
	ethClient, err := ethclient.Dial(clientUrl)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}
	defer ethClient.Close()

	// 创建合约对象
	if err != nil {
		log.Fatalf("Failed to instantiate a Contract: %v", err)
	}

	//封装过滤查询条件
	query := ethereum.FilterQuery{
		Addresses: []common.Address{common.HexToAddress(conAddr)},
		Topics:    [][]common.Hash{{crypto.Keccak256Hash([]byte("MyEvent(uint256)"))}},
	}

	// 创建事件订阅
	logsChan := make(chan types.Log) // 用于接受事件的通道
	sub, err := ethClient.SubscribeFilterLogs(context.Background(), query, logsChan)
	if err != nil {
		log.Fatalf("Failed to subscribe to events: %v", err)
	}
	// 处理订阅事件
	for {
		select {
		case err := <-sub.Err(): // 订阅失败的错误
			log.Fatalf("Subscription failed: %v", err)
		case vLog := <-logsChan: // 接收事件
			fmt.Printf("Received log: %+v\n", vLog)
			return
		}
	}
}

*/
