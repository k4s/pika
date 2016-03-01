package broker

import (
	"strings"
)

func NewBrokerClient(addr string) (Client interface{}) {
	if strings.Contains(addr, "redis") {
		Client = NewPikaRedisClient(addr)
	} else if strings.Contains(addr, "amqp") {
		Client = NewPikaRabbitMQClient(addr)
	} else {
		Client = nil
	}

	return

}
