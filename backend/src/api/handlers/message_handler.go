package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"wrapup/database"
	"wrapup/models"
	"wrapup/services"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MessageHandler struct {
	services.UserService
	services.MessageService
}

func NewMessageHandler(db *database.Client) *MessageHandler {
	return &MessageHandler{UserService: &services.UserServiceImpl{Db: db}, MessageService: &services.MessageServiceImpl{Db: db}}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var connections = make(map[string]*websocket.Conn)

func (mh *MessageHandler) HandleMessages(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()
	log.Println("WE ARE HERE 000")
	userID, _ := primitive.ObjectIDFromHex(c.Get("userID").(string))
	sender, err := mh.UserService.GetUser(&userID)

	if err != nil {
		log.Println("WE ARE HERE 001")
		return err
	}

	// send user online status to connected friends
	for connUserID, conn := range connections {
		connUserIDObjectID := ParseToPrimitive(connUserID)
		if mh.UserService.AreFriends(&sender.ID, &connUserIDObjectID) && conn != nil {
			sendUserStatusMessage(conn, userID, "online")
		}
		log.Println("WE ARE HERE 002")
	}

	connections[sender.ID.Hex()] = ws

	// Notify the user when one of their friends is disconnected
    defer func() {
        for connUserID, conn := range connections {
            connUserIDObjectID := ParseToPrimitive(connUserID)
            if mh.UserService.AreFriends(&sender.ID, &connUserIDObjectID) && conn != nil {
                sendUserStatusMessage(conn, userID, "offline")
            }
        }
        delete(connections, sender.ID.Hex())
    }()

	for {
		// Read
		var message models.Message
		_, msg, err := ws.ReadMessage()
		err = json.Unmarshal(msg, &message)
		if err != nil {
			sendErrorMessage(ws, &sender.ID, "error parsing the message")
			continue
		}
		log.Println("THE RECIPIENT ", message.RecipientID)
		_, err = mh.UserService.GetUser(&message.RecipientID)

		if err != nil {
			log.Println("WE ARE HERE 003", err)
			sendErrorMessage(ws, &sender.ID, "recipient is not a valid user")
			continue
		}
		message.SenderID = sender.ID
		if mh.UserService.AreFriends(&message.SenderID, &message.RecipientID) {
			recipientConn := connections[message.RecipientID.Hex()]
			mh.MessageService.CreateMessage(&message)
			if recipientConn != nil {
				sendPersonalMessage(recipientConn, &message)
			}
		}
	}
}

func sendErrorMessage(conn *websocket.Conn, receiver *primitive.ObjectID, text string) {
	conn.WriteJSON(map[string]interface{}{
		"type":    "error",
		"message": map[string]interface{}  {
			"user_id":  receiver,
			"message": text,
		},
	})
}

func sendPersonalMessage(conn *websocket.Conn, message *models.Message) {
	conn.WriteJSON(map[string]interface{}{
		"type":    "personal_message",
		"message": message,
	})
}

func sendUserStatusMessage(conn *websocket.Conn, userID primitive.ObjectID, status string) {
	conn.WriteJSON(map[string]interface{}{
		"type": "user_status",
		"message": map[string]interface{} {
			"user_id": userID,
			"status": status,
		},
	})
}

