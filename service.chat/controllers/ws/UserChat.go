package ChatControllerWS

import (
	"avirtan/work_craft/config"
	"fmt"

	"github.com/go-redis/redis/v7"
)

// ключи для redis
const (
	usersKey       = "users"
	userChannelFmt = "user:%s:channels"
	ChannelKey     = "channel:%s"
	ChatKey        = "chat:%s"
)

// пользователь для работы с redis, для чата
type UserChat struct {
	Name            string
	ChatName        string
	ChannelsHandler *redis.PubSub

	StopListenerChan chan struct{}
	Listening        bool

	//MessageChan chan redis.Message
}

// метод на создание пользователя
func CreateUser(name string, nameChat string) (*UserChat, error) {
	nameChat = fmt.Sprintf(ChatKey, nameChat)
	if _, err := config.DB_Redis.SAdd(nameChat, name).Result(); err != nil {
		fmt.Println("Ошибка подключения к чату")
		return nil, err
	}
	u := &UserChat{
		Name:             name,
		ChatName:         nameChat,
		StopListenerChan: make(chan struct{}),
	}
	return u, nil
}
