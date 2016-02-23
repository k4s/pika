package pika

import (
	"fmt"
	"strings"

	"github.com/k4s/pika/broker"
	"github.com/k4s/pika/tasks"
)

func PikaRun() {
	topicRun()
	plantRun()

}

//receive worker from flag
func workerlist() []string {
	workerlist := strings.Split(flWorker, ",")
	return workerlist
}

//receive publisher/subscriber's topic from flag
func topiclist() []string {
	topiclist := strings.Split(flTopic, ",")
	return topiclist
}

//receive producer/consumer's plant from flag
func plantlist() []string {
	plantlist := strings.Split(flPlant, ",")
	return plantlist
}

//running the publisher/subscriber's worker
func topicRun() {

	for _, topic := range topiclist() {
		redisClient := broker.NewPikaRedisClient(flBroker)
		pubsub, err := redisClient.Subscribe(topic)
		if err != nil {
			fmt.Println(err)
			return
		}
		go func() {
			defer pubsub.Close()
			for {
				msg, err := pubsub.ReceiveMessage()
				if err != nil {
					panic(err)
				}

				//fmt.Println(msg.Channel, msg.Payload)
				for _, worker := range workerlist() {
					tasks.TaskMap[worker](msg.Payload)

				}
			}
		}()

	}
}

//running the producer/consumer's worker
func plantRun() {
	for _, plant := range plantlist() {
		redisClient := broker.NewPikaRedisClient(flBroker)
		go func() {
			for {
				msg, err := redisClient.RPop(plant)
				if err != nil {
					continue
				}
				for _, worker := range workerlist() {
					tasks.TaskMap[worker](msg)
				}
			}
		}()
	}
}
