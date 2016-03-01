package pika

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/k4s/httprouter"
	"github.com/k4s/pika/broker"
	"github.com/k4s/pika/log"
)

func topicAdd(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	topicName := params.Get(":topicname")
	data, _ := ioutil.ReadAll(r.Body)
	//new the redis client
	if flBroker == "" {
		fmt.Fprintf(w, "no borker,you can use redis and rabbitmq")
		log.Logger("no borker,you can use redis and rabbitmq")
		return
	} else {
		client := broker.NewBrokerClient(flBroker)
		switch client := client.(type) {
		case *broker.PikaRedisClient:
			err := client.Publish(topicName, string(data))
			if err != nil {
				fmt.Fprintf(w, "%s publish err , msg:%s,err:%s", topicName, data, err)
				log.Logger(topicName, "publish err , msg:", data, ",err:", err)
				return
			}
		case *broker.PikaRabbitMQClient:
			err := client.Publish(topicName, string(data))
			if err != nil {
				fmt.Fprintf(w, "%s publish err , msg:%s,err:%s", topicName, data, err)
				log.Logger(topicName, "publish err , msg:", data, ",err:", err)
				return
			}
		}

	}

	fmt.Fprintf(w, "topic:%s publish succeed!", topicName)
}

func directAdd(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	directName := params.Get(":directname")
	data, _ := ioutil.ReadAll(r.Body)
	//new the redis client
	if flBroker == "" {
		fmt.Fprintf(w, "no borker,you can use redis and rabbitmq")
		log.Logger("no borker,you can use redis and rabbitmq")
		return
	} else {
		client := broker.NewBrokerClient(flBroker)
		switch client := client.(type) {
		case *broker.PikaRedisClient:
			err := client.LPush(directName, string(data))
			if err != nil {
				fmt.Fprintf(w, "%s producer err , msg:%s,err:%s", directName, data, err)
				log.Logger(directName, "publish err , msg:", data, ",err:", err)
				return
			}
		case *broker.PikaRabbitMQClient:
			fmt.Println(directName, string(data))
			err := client.Push(directName, string(data))
			if err != nil {
				fmt.Fprintf(w, "%s producer err , msg:%s,err:%s", directName, data, err)
				log.Logger(directName, "publish err , msg:", data, ",err:", err)
				return
			}
		}

	}

	fmt.Fprintf(w, "direct:%s producer values succeed!", directName)
}

func WebRun() {
	mux := httprouter.New()
	mux.Post("/topic(topic)/:topicname", topicAdd)
	mux.Post("/direct(direct)/:directname", directAdd)

	//  http.Handle("/", mux)

	if flHttp != "" {
		fmt.Println("restfulAPI run in http", flHttp)
		if err := http.ListenAndServe(flHttp, mux); err != nil {
			fmt.Println("http err:", err)
			log.Logger("no borker,you can use redis and rabbitmq")
		}
	}

}
