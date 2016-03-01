package pika

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/k4s/pika/broker"
	"github.com/k4s/pika/tasks"
)

func PikaRun() {
	if flTopic != "" {
		fmt.Println("Topic worker start....")
		topicRun()
	}
	if flDirect != "" {
		fmt.Println("Direct worker start....")
		directRun()
	}
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

//receive producer/consumer's direct from flag
func directlist() []string {
	directlist := strings.Split(flDirect, ",")
	return directlist
}

//running the publisher/subscriber's worker
func topicRun() {

	for _, topic := range topiclist() {
		client := broker.NewBrokerClient(flBroker)
		switch client := client.(type) {
		case *broker.PikaRedisClient:
			//			redisClient := broker.NewPikaRedisClient(flBroker)
			pubsub, err := client.Subscribe(topic)
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
						if worker == "" {
							continue
						}
						tasks.TaskMap[worker](msg.Payload)

					}
				}
			}()
		case *broker.PikaRabbitMQClient:
			pubsub, err := client.Subscribe(topic)
			if err != nil {
				return
			}
			go func() {
				for {
					for m := range pubsub {
						s := bytes.NewBuffer(m.Body)
						for _, worker := range workerlist() {
							if worker == "" {
								continue
							}
							tasks.TaskMap[worker](s.String())

						}
					}
				}
			}()

		}

	}
}

//running the producer/consumer's worker
func directRun() {
	for _, direct := range directlist() {
		client := broker.NewBrokerClient(flBroker)
		switch client := client.(type) {
		case *broker.PikaRedisClient:
			go func() {
				for {
					msg, err := client.RPop(direct)
					if err != nil {
						continue
					}
					for _, worker := range workerlist() {
						if worker == "" {
							continue
						}
						tasks.TaskMap[worker](msg)
					}
				}
			}()
		case *broker.PikaRabbitMQClient:
			go func() {
				for {
					msg, err := client.Pop(direct)
					if err != nil {
						continue
					}
					for _, worker := range workerlist() {
						if worker == "" {
							continue
						}
						tasks.TaskMap[worker](msg)
					}
				}
			}()
		}

	}
}
