
[中文READNE](https://github.com/k4s/pika/blob/master/readme_cn.md)

##Introduce:
pika is a MQ (message queue) written in Go and broker by Redis,RabbitMQ.<br/>

like python celery,but it is simple and easy to use. <br/>

with restfulAPI add async tasks.<br/>


##Apply:

###building a **redis** restfulAPI server:
> go run run.go -broker="redis://127.0.0.1:6379/0" -http="127.0.0.1:7778"

###client Demo:

publisher/subscriber by "kas", producer/consumer by "me" 
> go run run.go -worker=do -topic="kas" -direct="me" -broker="redis://127.0.0.1:6379/0"


publisher/subscriber by "kas"
> go run run.go -worker=add -topic="kas" -broker="redis://127.0.0.1:6379/0"



###building a **rabbitmq** restfulAPI server:
> go run run.go -broker="amqp://guest:guest@localhost:5672/" -http="127.0.0.1:7778"

##client Demo:

publisher/subscriber by "kas", producer/consumer by "me"
> go run run.go -worker=do -topic="kas" -direct="me" -broker="amqp://guest:guest@localhost:5672/"

publisher/subscriber by "kas"
> go run run.go -worker=add -topic="kas" -broker="amqp://guest:guest@localhost:5672/"