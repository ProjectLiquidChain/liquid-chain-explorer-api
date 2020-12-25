package surf

import (
	"encoding/hex"
	"encoding/json"
	"math"
	"net/http"

	"github.com/QuoineFinancial/liquid-chain-explorer-api/database"
	"github.com/QuoineFinancial/liquid-chain-explorer-api/node"
)

// GetTxsParams is params to GetAccount transaction
type GetTxsParams struct {
	paginationParams
}

// GetTxsResult is result of GetAccount
type GetTxsResult struct {
	paginationResult
	Transactions []node.Transaction `json:"transactions"`
	Receipts     []node.Receipt     `json:"receipts"`
}

// GetTxs lookup txs for an account
func (service Service) GetTxs(r *http.Request, params *GetTxsParams, result *GetTxsResult) error {
	limit := params.Limit
	if limit == 0 {
		limit = defaultLimit
	}

	var txs []database.Transaction
	if err := service.db.
		Order("transactions.id DESC").
		Offset(limit * params.Page).
		Limit(limit).
		Find(&txs).Error; err != nil {
		return err
	}

	nodeTxs := []node.Transaction{}
	receipts := []node.Receipt{}

	for _, tx := range txs {
		hashByte, err := hex.DecodeString(tx.Hash)
		if err != nil {
			return err
		}

		// Get tx
		nodeTxByte, err := service.txStorage.Get(hashByte)
		if err != nil {
			return err
		}

		var nodeTx node.Transaction
		if err := json.Unmarshal(nodeTxByte, &nodeTx); err != nil {
			return err
		}

		nodeTxs = append(nodeTxs, nodeTx)

		// Get receipt
		receiptByte, err := service.receiptStorage.Get(hashByte)
		if err != nil {
			return err
		}

		var receipt node.Receipt
		if err := json.Unmarshal(receiptByte, &receipt); err != nil {
			return err
		}

		receipts = append(receipts, receipt)
	}

	result.Transactions = nodeTxs
	result.Receipts = receipts

	var count int64
	if err := service.db.
		Model(&database.Transaction{}).
		Count(&count).Error; err != nil {
		return err
	}
	result.TotalPages = int(math.Ceil(float64(count) / float64(limit)))

	return nil
}
