package main

import (
	"encoding/json"
	"github.com/MahmoudMekki/MDS-task/clients/rabbitMQ"
	"github.com/MahmoudMekki/MDS-task/database"
	"github.com/MahmoudMekki/MDS-task/pkg/models"
	"github.com/MahmoudMekki/MDS-task/pkg/repo/orderDAL"
	"github.com/MahmoudMekki/MDS-task/pkg/repo/productsDAL"
	"github.com/rs/zerolog/log"
)

func init() {
	err := database.CreateDBConnection()
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	rabbitMQ.GetRabbitMQCConsumeChannel()
}
func main() {
	channel := rabbitMQ.GetRabbitMQCConsumeChannel()
	messages, err := channel.Consume(
		models.OrdersMQTopic,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal().Msg(err.Error())
	}
	forever := make(chan bool)
	log.Info().Msg("Queue worker is running ")
	go func() {
		for message := range messages {
			var records []models.OrderCSV
			err := json.Unmarshal(message.Body, &records)
			if err != nil {
				log.Err(err).Msg(err.Error())
				continue
			}
			for _, v := range records {
				prod, existed, err := productsDAL.GetProdBySkuAndCountry(v.SKU, v.Country)
				if err != nil {
					log.Err(err).Msg(err.Error())
					continue
				}
				if !existed {
					var amount int
					if v.StockChange > 0 {
						amount = v.StockChange
					}
					prod = models.Product{
						SKU:         v.SKU,
						CountryCode: v.Country,
						Name:        v.Name,
						Amount:      amount,
					}
					prod, err = productsDAL.CreateProduct(prod)
					if err != nil {
						log.Err(err).Msg(err.Error())
						continue
					}
					order := models.Order{
						ProductId: prod.Id,
						Amount:    amount,
					}
					order, err = orderDAL.CreateOrder(order)
					if err != nil {
						log.Err(err).Msg(err.Error())
						continue
					}
					continue
				}
				if prod.Amount+v.StockChange < 0 {
					log.Info().Msg("insufficient stock balance")
					continue
				}
				err = productsDAL.UpdateProdStock(prod.Id, v.StockChange)
				if err != nil {
					log.Err(err).Msg(err.Error())
					continue
				}
				order := models.Order{
					ProductId: prod.Id,
					Amount:    v.StockChange,
				}
				order, err = orderDAL.CreateOrder(order)
				if err != nil {
					log.Err(err).Msg(err.Error())
				}
			}
		}
	}()

	<-forever
}
