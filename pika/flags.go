package pika

import (
	"flag"
)

var (
	flBroker string
	flWorker string
	flTopic  string
	flDirect string
	flHttp   string
	flTcp    string
)

func Flags() {
	flag.StringVar(&flBroker, "broker", "", "redis,rabbitmq Address or other broker. eg:redis://password@127.0.0.1:6379/0")
	flag.StringVar(&flWorker, "worker", "", "the name of Tasksworker")
	flag.StringVar(&flTopic, "topic", "", "topic's name for publish/subscribe")
	flag.StringVar(&flDirect, "direct", "", "direct's name for produce/consume")
	flag.StringVar(&flHttp, "http", "", "http address and port. eg:127.0.0.1:7778")
	flag.StringVar(&flTcp, "tcp", "127.0.0.1:7779", "tcp address and port")
	flag.Parse()
}
