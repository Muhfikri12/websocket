package handler

import (
	"encoding/json"
	"fmt"
	"project/helper"
	"project/service"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type WebSocketController struct {
	service service.Service
	logger  *zap.Logger
}

func NewWebSocketController(service service.Service, logger *zap.Logger) *WebSocketController {
	return &WebSocketController{service: service, logger: logger}
}

func (wsc *WebSocketController) DashboardGetAllOrder(c *gin.Context) {
	conn, err := helper.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Failed to upgrade connection:", err)
		return
	}
	defer conn.Close()

	fmt.Println("Client connected")

	// Kirim data real-time ke klien
	for {

		data, err := wsc.service.Order.GetAllOrder()
		if err != nil {
			wsc.logger.Error("error get data")
			return
		}

		jsonData, err := json.Marshal(data)
		if err != nil {
			wsc.logger.Error("Error marshalling data to JSON", zap.Error(err))
			return
		}

		// Kirim data ke klien
		err = conn.WriteMessage(websocket.TextMessage, jsonData)
		if err != nil {
			fmt.Println("Error writing message:", err)
			break
		}

		// Delay sebelum mengirim data berikutnya
		time.Sleep(1 * time.Second)
	}
}
