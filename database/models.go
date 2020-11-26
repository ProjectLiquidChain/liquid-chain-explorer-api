package database

// Account represents an account of liquid-chain
type Account struct {
	ID      uint   `gorm:"primarykey"`
	Address string `gorm:"index"`
}

// Asset represents a token asset of Account {
type Asset struct {
	ID        uint `gorm:"primarykey"`
	TokenID   uint `gorm:"index"`
	AccountID uint `gorm:"index"`
	Balance   string
	Token     Token   `gorm:"foreignKey:TokenID"`
	Account   Account `gorm:"foreignKey:AccountID"`
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
	ID         uint    `gorm:"primarykey" json:"-"`
	Block      uint64  `gorm:"index" json:"block"`
	SenderID   uint    `gorm:"index" json:"-"`
	ReceiverID uint    `gorm:"index" json:"-"`
	Hash       string  `gorm:"uniqueIndex" json:"hash"`
	Sender     Account `gorm:"foreignKey:SenderID" json:"sender"`
	Receiver   Account `gorm:"foreignKey:ReceiverID" json:"receiver"`
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

	Transaction Transaction `gorm:"foreignKey:TransactionID" json:"-"`
	FromAccount Account     `gorm:"foreignKey:FromAccountID" json:"from"`
	ToAccount   Account     `gorm:"foreignKey:ToAccountID" json:"to"`
	Token       Token       `gorm:"foreignKey:TokenID" json:"token"`
}
