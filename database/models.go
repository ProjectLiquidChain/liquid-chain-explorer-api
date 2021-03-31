package database

// Block represents a block in liquid-chain
type Block struct {
	ID     uint   `gorm:"primaryKey"`
	Height uint   `gorm:"index"`
	Time   uint64 `gorm:"index"`
	Hash   string
}

// Account represents an account of liquid-chain
type Account struct {
	ID      uint   `gorm:"primarykey" json:"-"`
	Address string `gorm:"index" json:"address"`
}

// Asset represents a token asset of Account {
type Asset struct {
	ID        uint    `gorm:"primarykey" json:"-"`
	TokenID   uint    `gorm:"index" json:"-"`
	AccountID uint    `gorm:"index" json:"-"`
	Balance   string  `json:"balance"`
	Token     Token   `gorm:"foreignKey:TokenID" json:"token"`
	Account   Account `gorm:"foreignKey:AccountID" json:"-"`
}

// Token represents a token of liquid-chain
type Token struct {
	ID       uint   `gorm:"primarykey" json:"-"`
	Address  string `gorm:"index" json:"address"`
	Currency string `gorm:"index" json:"currency"`
	Decimals int    `json:"decimals"`
}

// Transaction represents tx of liquid-chain
type Transaction struct {
	ID         uint     `gorm:"primarykey" json:"-"`
	Block      uint64   `gorm:"index" json:"block"`
	SenderID   uint     `gorm:"index" json:"-"`
	ReceiverID uint     `gorm:"index" json:"-"`
	Hash       string   `gorm:"uniqueIndex" json:"hash"`
	Sender     *Account `gorm:"foreignKey:SenderID" json:"sender,omitempty"`
	Receiver   *Account `gorm:"foreignKey:ReceiverID" json:"receiver,omitempty"`
}

// Transfer represents Transfer event
type Transfer struct {
	ID            uint   `gorm:"primaryKey" json:"-"`
	EventIndex    int    `json:"index"`
	TransactionID uint   `gorm:"index" json:"-"`
	FromAccountID uint   `gorm:"index" json:"-"`
	ToAccountID   uint   `gorm:"index" json:"-"`
	TokenID       uint   `gorm:"index" json:"-"`
	Amount        string `json:"amount"`
	Memo          string `json:"memo"`

	Transaction Transaction `gorm:"foreignKey:TransactionID" json:"transaction"`
	FromAccount Account     `gorm:"foreignKey:FromAccountID" json:"from"`
	ToAccount   Account     `gorm:"foreignKey:ToAccountID" json:"to"`
	Token       Token       `gorm:"foreignKey:TokenID" json:"token"`
}
