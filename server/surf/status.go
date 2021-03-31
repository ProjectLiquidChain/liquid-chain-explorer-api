package surf

import (
	"encoding/hex"
	"encoding/json"
	"math"
	"net/http"

	"github.com/QuoineFinancial/liquid-chain-explorer-api/database"
	"github.com/QuoineFinancial/liquid-chain-explorer-api/node"
)

// GetStatusParams is params to GetAccount transaction
type GetStatusParams struct {
}

// GetStatusResult is result of GetAccount
type GetStatusResult struct {
	AverageBlockTime  float64 `json:"averageBlockTime"`
	TotalTransactions int     `json:"totalTxs"`
	Price             string  `json:"price"`
}

func (service Service) getBlock(hash string) (*node.Block, error) {
	hashByte, err := hex.DecodeString(hash)
	if err != nil {
		return nil, err
	}

	blockByte, err := service.blockStorage.Get(hashByte)
	if err != nil {
		return nil, err
	}

	var block node.Block
	if err := json.Unmarshal(blockByte, &block); err != nil {
		return nil, err
	}

	return &block, nil
}

func (service Service) getTotalTxs() (int, error) {
	var count int64
	if err := service.db.
		Model(&database.Transaction{}).
		Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}

func (service Service) getPrice() (string, error) {
	const productID = "57"
	product, err := service.liquidAPI.GetProduct(productID)
	if err != nil {
		return "", err
	}
	return product.LastTradedPrice, nil
}

func (service Service) getAverageBlockTime() (float64, error) {
	var firstBlock, lastBlock database.Block

	if err := service.db.Where(database.Block{Height: 0}).First(&firstBlock).Error; err != nil {
		return 0, err
	}
	firstNodeBlock, err := service.getBlock(firstBlock.Hash)
	if err != nil {
		return 0, err
	}

	if err := service.db.Order("height DESC").First(&lastBlock).Error; err != nil {
		return 0, err
	}
	lastNodeBlock, err := service.getBlock(lastBlock.Hash)
	if err != nil {
		return 0, err
	}

	return math.Round(10*float64(lastNodeBlock.Time-firstNodeBlock.Time)/float64(lastNodeBlock.Height)) / 10, nil
}

// GetStatus lookup txs for an account
func (service Service) GetStatus(r *http.Request, params *GetStatusParams, result *GetStatusResult) error {
	duration, err := service.getAverageBlockTime()
	if err != nil {
		return err
	}

	totalTxs, err := service.getTotalTxs()
	if err != nil {
		return err
	}

	price, err := service.getPrice()
	if err != nil {
		return err
	}

	result.AverageBlockTime = duration
	result.TotalTransactions = totalTxs
	result.Price = price

	return nil
}
