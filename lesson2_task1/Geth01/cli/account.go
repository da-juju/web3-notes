package cli

import (
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
)

// NewAccount 新建账户
func NewAccount(cli *rpc.Client, pass string) (string, error) {
	var res string
	err := cli.Call(&res, "personal_newAccount", pass)
	if err != nil {
		fmt.Println("[NewAccount Err]", err)
		return "", err
	}
	return res, nil
}

// ListAccount 查询所有账户
func ListAccount(cli *rpc.Client) ([]string, error) {
	var res []string
	err := cli.Call(&res, "personal_listAccounts")
	if err != nil {
		fmt.Println("[ListAccount Err]", err)
		return nil, err
	}
	return res, nil
}
