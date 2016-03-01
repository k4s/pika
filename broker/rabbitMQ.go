package broker

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/k4s/pika/common/util"
	"github.com/streadway/amqp"
)

func newRabbitMQClient(addr string) *amqp.Connection {
	conn, err := amqp.Dial(addr) //amqp://guest:guest@localhost:5672/
	if err != nil {
		panic(err)
	}
	return conn
}

type PikaRabbitMQClient struct {
	client      *amqp.Connection
	exchangeMap map[string]bool
}

func NewPikaRabbitMQClient(rabbitMQAddr string) *PikaRabbitMQClient {
	//get db from rabbitMQ

	return &PikaRabbitMQClient{
		client:      newRabbitMQClient(rabbitMQAddr),
		exchangeMap: make(map[string]bool),
	}

}
func (p *PikaRabbitMQClient) newExchange(topic string) error {
	if p.exchangeMap[topic] == true {
		return errors.New("topicName of exchange is exist")
	}
	ch, err := p.client.Channel()
	defer ch.Close()
	if err != nil {
		return err
	}
	err = ch.ExchangeDeclare(topic, "fanout", false, false, false, false, nil)
	if err != nil {
		return err
	}
	p.exchangeMap[topic] = true
	return nil
}
func (p *PikaRabbitMQClient) Publish(topic string, pmsg string) error {
	if p.exchangeMap[topic] == false {
		err := p.newExchange(topic)
		if err != nil {
			return err
		}
	}
	ch, err := p.client.Channel()
	defer ch.Close()
	err = ch.Publish(
		topic, // exchange
		"",    // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(pmsg),
		})
	return err
}

func (p *PikaRabbitMQClient) Subscribe(topic string) (<-chan amqp.Delivery, error) {
	if p.exchangeMap[topic] == false {
		err := p.newExchange(topic)
		if err != nil {
			return nil, err
		}
	}
	ch, err := p.client.Channel()
	if err != nil {
		return nil, err
	}
	//new a queue to bind on topic exchange
	queueName := "pika_" + util.RandSeq(int(10))
	_, err = ch.QueueInspect(queueName)
	//if rabbitmq doesn't have the queue of queueName,new a queue of queueName
	if err != nil {
		//if err != nil,the channel will close,so must open Once again
		ch, err = p.client.Channel()
		_, err := ch.QueueDeclare(
			queueName, // name
			true,      // durable
			true,      // delete when unused
			false,     // exclusive
			false,     // no-wait
			nil,       // arguments
		)
		if err != nil {
			fmt.Println("QueueDeclare error:", err)
			return nil, err
		}

	}
	err = ch.QueueBind(queueName, "", topic, false, nil)
	if err != nil {
		return nil, err
	}
	//QueueBind("pagers", "alert", "log", false, nil)
	//QueueBind("emails", "info", "log", false, nil)
	//Delivery       Exchange  Key       Queue
	//-----------------------------------------------
	//key: alert --> log ----> alert --> pagers
	//key: info ---> log ----> info ---> emails
	msgs, err := ch.Consume(queueName, "", true, false, false, false, nil)
	return msgs, err
	//	waitMsg := make(chan string)
	//	go func() {
	//		for m := range msgs {
	//			s := bytes.NewBuffer(m.Body)
	//			waitMsg <- s.String()
	//		}
	//	}()
	//	msg <- waitMsg
	//	return msg, err
}

func (p *PikaRabbitMQClient) Push(queueName string, values ...string) error {
	ch, err := p.client.Channel()
	defer ch.Close()
	if err != nil {
		return err
	}
	_, err = ch.QueueInspect(queueName)
	//if rabbitmq doesn't have the queue of queueName,new a queue of queueName
	if err != nil {
		//if err != nil,the channel will close,so must open Once again
		ch, err = p.client.Channel()
		_, err := ch.QueueDeclare(
			queueName, // name
			true,      // durable
			false,     // delete when unused
			false,     // exclusive
			false,     // no-wait
			nil,       // arguments
		)
		return err
	}

	for _, value := range values {
		err = ch.Publish(
			"",        // exchange
			queueName, // routing key
			false,     // mandatory
			false,     // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(value),
			})
	}

	return err
}

func (p *PikaRabbitMQClient) Pop(queueName string) (string, error) {
	ch, err := p.client.Channel()
	if err != nil {
		return "", err
	}
	_, err = ch.QueueInspect(queueName)
	if err != nil {
		return "", err
	}
	msgs, err := ch.Consume(queueName, "", true, false, false, false, nil)
	if err != nil {
		return "", err
	}
	waitMsg := make(chan string)
	go func() {
		for m := range msgs {
			s := bytes.NewBuffer(m.Body)
			waitMsg <- s.String()
		}
	}()
	msg := <-waitMsg
	return msg, err
}
