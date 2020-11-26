package node

type Argument struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Call struct {
	Contract string     `json:"contract,omitempty"`
	Name     string     `json:"name,omitempty"`
	Args     []Argument `json:"args,omitempty"`
}

type Receipt struct {
	Index       uint32 `json:"index"`
	Transaction string `json:"transaction"`
	Result      string `json:"result"`
	GasUsed     uint32 `json:"gasUsed"`
	Code        byte   `json:"code"`
	Events      []Call `json:"events"`
	PostState   string `json:"postState"`
}

type TransactionType string

const (
	transactionTypeDeploy TransactionType = "deploy"
	transactionTypeInvoke TransactionType = "invoke"
)

type Transaction struct {
	Hash        string          `json:"hash"`
	Type        TransactionType `json:"type"`
	BlockHeight uint64          `json:"height"`
	Version     uint16          `json:"version"`
	Sender      string          `json:"sender"`
	Nonce       uint64          `json:"nonce"`
	Receiver    string          `json:"receiver"`
	Payload     Call            `json:"payload"`
	GasPrice    uint32          `json:"gasPrice"`
	GasLimit    uint32          `json:"gasLimit"`
	Signature   []byte          `json:"signature"`
}

type Block struct {
	Hash            string        `json:"hash"`
	Transactions    []Transaction `json:"transactions"`
	Receipts        []Receipt     `json:"receipts"`
	Height          uint64        `json:"height"`
	Time            uint64        `json:"time"`
	Parent          string        `json:"parent"`
	StateRoot       string        `json:"stateRoot"`
	TransactionRoot string        `json:"transactionRoot"`
	ReceiptRoot     string        `json:"receiptRoot"`
}
