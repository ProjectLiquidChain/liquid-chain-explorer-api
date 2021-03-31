package main

import (
	"log"
	"os"

	"github.com/QuoineFinancial/liquid-chain-explorer-api/database"
	"github.com/QuoineFinancial/liquid-chain-explorer-api/node"
	"github.com/sirupsen/logrus"
)

func copyBlockTimeFromStorageToDatabase() error {
	db := database.New(os.Getenv("DATABASE_URL"))
	db.AutoMigrate(database.Block{})
	node := node.New(os.Getenv("NODE_URL"))

	var blocks []database.Block
	db.Model(&database.Block{}).Find(&blocks).Where("time IS NULL")
	for _, block := range blocks {
		nodeBlock, err := node.GetBlock(uint64(block.Height))
		if err != nil {
			panic(err)
		}
		logrus.Infof("Update time for block %d -> %d", block.Height, nodeBlock.Time)

		if err := db.Model(&block).Updates(database.Block{Time: nodeBlock.Time}).Error; err != nil {
			return err
		}
	}

	return nil
}

func main() {
	if err := copyBlockTimeFromStorageToDatabase(); err != nil {
		log.Println(err)
	}
}
