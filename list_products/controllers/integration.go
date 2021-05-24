package controllers

import (
	"encoding/json"
	"fmt"
	"list_product/config"
	model "list_product/models"
)

// SetupRouter build gin engine
func (server *Server) PoolingSqs() {
	for true {
		messageFromQueue, err := server.SQSMessager.ReadMessages(config.ENV.SQSQueueURL, 1)

		if err != nil {
			fmt.Println("Error to read message from queue")
			continue
		}

		for _, message := range messageFromQueue {
			messageModel := model.User{}
			err = json.Unmarshal([]byte(*message.Body), &messageModel)
			if err != nil {
				fmt.Println("Error to unmarshal message")
				continue
			}

			var listProduct model.ListProducts
			listProduct.UserId = messageModel.ID

			listProductCreated, err := listProduct.SaveListProduct(server.DB)
			if err != nil {
				fmt.Println("Error to create list of products")
				continue
			}

			fmt.Println(listProductCreated)

			server.SQSMessager.DeleteMessage(message, config.ENV.SQSQueueURL)
		}
	}
}
