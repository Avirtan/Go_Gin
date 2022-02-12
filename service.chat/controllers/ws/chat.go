package ChatControllerWS

import (
	"avirtan/work_craft/config"
	jwtHandler "avirtan/work_craft/pkg/jwt"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// поля для отправки сообщения в чат
type msg struct {
	Content string `json:"content,omitempty"`
}

//настройка вебсокета
var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Chat godoc
// @Summary 	Chat
// @Description connect to chat ws
// @Tags 		Chat
// @Param       name    query     string  true  "name chat"
// @Param       token   query     string  true  "token user"
// @Router 		/ws/chat/:name/:token [get]
// метод на обработку запроса к вебсокету и установке соединения,
// а так же запуск всех основных функций для взаимодействия с вебсокетом
func ChatWebSocketHandler(w http.ResponseWriter, r *http.Request, ctx *gin.Context) {
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Failed to set websocket upgrade: %+v\n", err)
		return
	}
	date, ok := jwtHandler.ValidateToken(ctx.Param("token"))
	if !ok {
		conn.WriteJSON(gin.H{"err": "invalid token"})
		conn.Close()
		return
	}
	user, err := onConnect(r, conn, date["login"].(string), ctx.Param("name"))
	if err != nil {
		fmt.Println("Ошибка подключения")
		conn.WriteJSON(gin.H{"err": err})
		return
	}
	closeCh := onDisconnect(r, conn, user)

	onChannelMessage(conn, r, user)

loop:
	for {
		select {
		case <-closeCh:
			break loop
		default:
			onUserMessage(conn, r, user)
		}
	}
}

// метод сробатывает когда пользователь подключается
func onConnect(r *http.Request, conn *websocket.Conn, username string, nameChat string) (*UserChat, error) {
	//fmt.Println("connected from:", conn.RemoteAddr(), "user:", username)
	u, err := CreateUser(username, nameChat)
	if err != nil {
		return u, err
	}
	return u, nil
}

//метод срабатывает при разрывае соединения
func onDisconnect(r *http.Request, conn *websocket.Conn, user *UserChat) chan struct{} {
	closeCh := make(chan struct{})
	conn.SetCloseHandler(func(code int, text string) error {
		fmt.Println("connection closed for user", user.Name)
		if user.ChannelsHandler != nil {
			if err := config.DB_Redis.SRem(user.ChatName, user.Name).Err(); err != nil {
				return err
			}
			if err := user.ChannelsHandler.Unsubscribe(); err != nil {
				return err
			}
			if err := user.ChannelsHandler.Close(); err != nil {
				return err
			}
		}
		if user.Listening {
			user.StopListenerChan <- struct{}{}
		}
		close(closeCh)
		return nil
	})

	return closeCh
}

// метод на отправку сообщений в чат
func onUserMessage(conn *websocket.Conn, r *http.Request, user *UserChat) {
	var massage msg
	err := conn.ReadJSON(&massage)
	if err != nil {
		//fmt.Println("ошибка чтения json")
		conn.WriteJSON(gin.H{"err": err.Error()})
		return
	}
	err = config.DB_Redis.Publish(user.ChatName, massage.Content).Err()
	if err != nil {
		conn.WriteJSON(gin.H{"err": err.Error()})
	}
}

// метод на обработку полученных сообщений от других пользователей
func onChannelMessage(conn *websocket.Conn, r *http.Request, user *UserChat) {
	pubSub := config.DB_Redis.Subscribe(user.ChatName)
	user.ChannelsHandler = pubSub
	go func() {
		user.Listening = true
		fmt.Println("starting the listener for user:", user.Name, "on channels:", user.ChatName)
		for {
			select {
			case msg, ok := <-pubSub.Channel():
				if !ok {
					return
				}
				if err := conn.WriteJSON(gin.H{"user": user.Name, "msg": msg.Payload}); err != nil {
					fmt.Println(err)
				}
			case <-user.StopListenerChan:
				//fmt.Println("stopping the listener for user:", user.Name)
				return
			}
		}
	}()
}
