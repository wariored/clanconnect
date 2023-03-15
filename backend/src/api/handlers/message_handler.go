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
	userID, _:= primitive.ObjectIDFromHex(c.Get("userID").(string))
	sender, err := mh.UserService.GetUser(&userID)

	if err != nil {
	log.Println("WE ARE HERE 001")
		return err
	}
	connections[sender.ID.Hex()] = ws
	
	// send user online status to connected friends
	for connUserID, conn := range connections {
		if mh.UserService.AreFriends(&sender.ID, connUserID) && conn != nil {
			conn.WriteJSON(map[string]interface{}{
			  "type":     "user_status",
			  "userID":   userID,
			  "isOnline": true,
			})
		}
		log.Println("WE ARE HERE 002")
	}

	for {
		// Read
		// Parse message
        var message models.Message
		_, msg, err := ws.ReadMessage()
        err = json.Unmarshal(msg, &message)
		if err != nil {
			sendErrorMessage(ws, &sender.ID, "error parsing the message")
		}
		log.Println("THE RECIPIENT ", message.RecipientID)	
		_, err = mh.UserService.GetUser(message.RecipientID)

		if err != nil {
			log.Println("WE ARE HERE 003", err)
			sendErrorMessage(ws, &sender.ID, "recipient is not a valid user")
			continue
		}
		message.SenderID = &sender.ID
		if mh.UserService.AreFriends(message.SenderID, message.RecipientID){
			recipientConn := connections[message.RecipientID]
			log.Println("ARE FRIENDS ", recipientConn)
			if recipientConn != nil {
				sendPersonalMessage(recipientConn, &message)
			}
			mh.MessageService.CreateMessage(&message)
		}	
	}
}

func sendErrorMessage(conn *websocket.Conn, receiver *primitive.ObjectID, text string){
	conn.WriteJSON(map[string]interface{}{
		  "type":     "error",
		  "userID":   receiver,
		  "message": text,
	})
}

func sendPersonalMessage(conn *websocket.Conn, message *models.Message){ 
	conn.WriteJSON(map[string]interface{}{
		  "type":     "personal_message",
		  "sender_id":   message.SenderID,
		  "recipient_id":   message.RecipientID,
		  "text": message.Text,
	})
}

