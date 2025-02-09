package chainlink

import (
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type ChainlinkClient struct {
	ContractAddress common.Address
	Instance        *SmartContract // Replace with your actual contract instance
}

func NewChainlinkClient(contractAddress string, backend bind.ContractBackend) (*ChainlinkClient, error) {
	address := common.HexToAddress(contractAddress)
	instance, err := NewSmartContract(address, backend)
	if err != nil {
		return nil, err
	}
	return &ChainlinkClient{
		ContractAddress: address,
		Instance:        instance,
	}, nil
}

func (c *ChainlinkClient) FetchPrice() (*big.Int, error) {
	price, err := c.Instance.GetPrice(nil) // Replace `GetPrice` with your actual contract method
	if err != nil {
		return nil, err
	}
	log.Printf("Fetched price: %s", price.String())
	return price, nil
}