package database

func (db Database) migrate() {
	db.AutoMigrate(Block{})
	db.AutoMigrate(Account{})
	db.AutoMigrate(Asset{})
	db.AutoMigrate(Token{})
	db.AutoMigrate(Transaction{})
	db.AutoMigrate(Transfer{})
}
