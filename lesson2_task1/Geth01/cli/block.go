package cli

import (
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
	"strconv"
)

// BlockNumber 获取区块高度
func BlockNumber(cli *rpc.Client) (int, error) {
	var res string
	err := cli.Call(&res, "eth_blockNumber")
	if err != nil {
		fmt.Println("[GetBlockNumber err]", err)
		return 0, err
	}
	parseInt, _ := strconv.ParseInt(res, 0, 0)
	return int(parseInt), nil
}
