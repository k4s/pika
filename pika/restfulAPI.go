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
		redisClient := broker.NewPikaRedisClient(flBroker)
		fmt.Println(topicName, string(data))
		err := redisClient.Publish(topicName, string(data))
		if err != nil {
			fmt.Fprintf(w, "%s publish err , msg:%s,err:%s", topicName, data, err)
			log.Logger(topicName, "publish err , msg:", data, ",err:", err)
			return
		}
	}

	fmt.Fprintf(w, "topic:%s publish succeed!", topicName)
}

func plantAdd(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	plantName := params.Get(":plantname")
	data, _ := ioutil.ReadAll(r.Body)
	//new the redis client
	if flBroker == "" {
		fmt.Fprintf(w, "no borker,you can use redis and rabbitmq")
		log.Logger("no borker,you can use redis and rabbitmq")
		return
	} else {
		redisClient := broker.NewPikaRedisClient(flBroker)
		fmt.Println(plantName, string(data))
		err := redisClient.LPush(plantName, string(data))
		if err != nil {
			fmt.Fprintf(w, "%s producer err , msg:%s,err:%s", plantName, data, err)
			log.Logger(plantName, "publish err , msg:", data, ",err:", err)
			return
		}
	}

	fmt.Fprintf(w, "plant:%s producer values succeed!", plantName)
}

func WebRun() {
	mux := httprouter.New()
	mux.Post("/topic(topic)/:topicname", topicAdd)
	mux.Post("/plant(plant)/:plantname", plantAdd)

	//  http.Handle("/", mux)

	if flHttp != "" {
		fmt.Println("restfulAPI run in http", flHttp)
		if err := http.ListenAndServe(flHttp, mux); err != nil {
			fmt.Println("http err:", err)
			log.Logger("no borker,you can use redis and rabbitmq")
		}
	}

}
