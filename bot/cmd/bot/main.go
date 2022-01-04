package main

import (
	"github.com/amrchnk/ozon_go_course/bot/internal/config"
	"github.com/amrchnk/ozon_go_course/bot/internal/repository"
	"github.com/amrchnk/ozon_go_course/bot/internal/repository/boltDB"
	telegram "github.com/amrchnk/ozon_go_course/bot/internal/service/telegram"
	"github.com/boltdb/bolt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}
	//log.Println("ENV:",cfg)
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Fatal(err)
	}
	bot.Debug = true


	db, err := initDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	productRepository:=boltDB.NewProductRepository(db)

	tgBot:=telegram.NewBot(bot,productRepository)

	err = tgBot.Start()
	if err != nil {
		log.Fatal(err)
	}
	/*go func(){

	}()*/
}

func initDB(cfg *config.Config) (*bolt.DB, error) {
	db, err := bolt.Open(cfg.DBPath, 0600, nil)
	if err != nil {
		return nil, err
	}

	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(repository.Product))
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}
	return db, nil
}
