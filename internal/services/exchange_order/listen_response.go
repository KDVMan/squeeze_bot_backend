package services_exchange_order

import (
	"log"
)

func (object *exchangeOrderServiceImplementation) listenResponse() {
	for message := range object.orderPlaceService.GetReadChannel() {
		log.Println("Order response:", string(message))

		// Пример парсинга (опционально адаптируй)
		// var response struct {
		// 	ClientOrderId string `json:"clientOrderId"`
		// 	Status        string `json:"status"`
		// }
		// if err := json.Unmarshal(msg, &response); err != nil {
		// 	log.Println("Failed to parse response:", err)
		// 	continue
		// }
		//
		// ws.mu.Lock()
		// bot, exists := ws.botMap[response.ClientOrderId]
		// ws.mu.Unlock()
		// if exists {
		// 	log.Println("Found bot:", bot.ID, "status:", response.Status)
		// 	// Тут можно обновить botModel (например, статус ордера)
		// }
	}
}
